package hs

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"github.com/bobacgo/admin-service/apps/ecode"
	"github.com/bobacgo/admin-service/pkg/kit/hs/response"
	"github.com/go-playground/validator/v10"
)

// HTTPMethodPrefix 定义HTTP方法前缀映射
var HTTPMethodPrefix = map[string]string{
	"Get":    http.MethodGet,
	"Post":   http.MethodPost,
	"Put":    http.MethodPut,
	"Delete": http.MethodDelete,
	"Patch":  http.MethodPatch,
}

// camelToKebab 将驼峰命名转换为斜杠分隔的路径
// 例如: GetUserInfo -> /user/info, GetList -> /list
func camelToKebab(s string) string {
	// 移除HTTP方法前缀
	for prefix := range HTTPMethodPrefix {
		if strings.HasPrefix(s, prefix) && len(s) > len(prefix) {
			s = s[len(prefix):]
			break
		}
	}

	if s == "" {
		return ""
	}

	// 将驼峰转换为斜杠分隔
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('/')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return "/" + result.String()
}

// extractHTTPMethod 从方法名前缀提取HTTP方法
func extractHTTPMethod(methodName string) string {
	for prefix, method := range HTTPMethodPrefix {
		if strings.HasPrefix(methodName, prefix) {
			return method
		}
	}
	return ""
}

// ServiceMethodInfo 包含服务方法的元数据
type ServiceMethodInfo struct {
	Method     reflect.Method
	HTTPMethod string
	Path       string
	MethodName string
}

// DiscoverServiceMethods 发现服务中符合自动路由规则的方法
// 规则:
// 1. 方法名以 Get/Post/Put/Delete/Patch 开头
// 2. 方法签名必须为: func(ctx context.Context, req *ReqType) (*RespType, error)
func DiscoverServiceMethods(service interface{}) []ServiceMethodInfo {
	var methods []ServiceMethodInfo
	serviceType := reflect.TypeOf(service)

	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)
		methodName := method.Name

		// 跳过未导出的方法
		if !method.IsExported() {
			continue
		}

		// 检查HTTP方法前缀
		httpMethod := extractHTTPMethod(methodName)
		if httpMethod == "" {
			continue
		}

		// 验证方法签名: func(ctx context.Context, req *ReqType) (*RespType, error)
		if !isValidServiceMethodSignature(method.Type) {
			slog.Warn("skip method with invalid signature", "method", methodName, "reason", "signature does not match func(ctx context.Context, req *ReqType) (*RespType, error)")
			continue
		}

		// 生成路由路径
		path := camelToKebab(methodName)

		methods = append(methods, ServiceMethodInfo{
			Method:     method,
			HTTPMethod: httpMethod,
			Path:       path,
			MethodName: methodName,
		})
	}

	return methods
}

// isValidServiceMethodSignature 检查方法签名是否符合自动路由要求
// 要求: func(ctx context.Context, req *ReqType) (*RespType, error)
// 即: 3个输入参数(receiver, context, request), 2个输出参数(response, error)
func isValidServiceMethodSignature(methodType reflect.Type) bool {
	// 检查输入参数数量: receiver + context + request = 3
	if methodType.NumIn() != 3 {
		return false
	}

	// 检查输出参数数量: response + error = 2
	if methodType.NumOut() != 2 {
		return false
	}

	// 检查第二个输入参数是否是 context.Context
	if !isContextType(methodType.In(1)) {
		return false
	}

	// 检查最后一个输出参数是否是 error
	if !isErrorType(methodType.Out(1)) {
		return false
	}

	// 检查第三个输入参数是否是指针类型 (request)
	reqType := methodType.In(2)
	if reqType.Kind() != reflect.Ptr {
		return false
	}

	// 检查第一个输出参数是否是指针类型 (response)
	respType := methodType.Out(0)
	if respType.Kind() != reflect.Ptr && respType.Kind() != reflect.Interface {
		// 允许接口类型(如interface{})作为返回值
		return false
	}

	return true
}

// isContextType 检查类型是否为 context.Context
func isContextType(t reflect.Type) bool {
	return t.String() == "context.Context"
}

// isErrorType 检查类型是否为 error
func isErrorType(t reflect.Type) bool {
	return t.String() == "error"
}

// HandlerConfig 用于处理程序的配置
type HandlerConfig struct {
	Validator *validator.Validate
}

// CreateAutoHandler 为服务方法创建HTTP处理程序
// 方法签名: func(ctx context.Context, req *ReqType) (*RespType, error)
func CreateAutoHandler(methodInfo ServiceMethodInfo, service interface{}, config *HandlerConfig) http.HandlerFunc {
	method := methodInfo.Method
	serviceValue := reflect.ValueOf(service)
	methodType := method.Type

	return func(w http.ResponseWriter, r *http.Request) {
		// 验证HTTP方法匹配
		if r.Method != methodInfo.HTTPMethod {
			response.JSON(w, response.Resp{
				Code: ecode.ErrCodeParam,
				Msg:  "method not allowed",
			})
			return
		}

		// 构建输入参数: [receiver, context, request]
		var args []reflect.Value
		args = append(args, serviceValue)
		args = append(args, reflect.ValueOf(r.Context()))

		// 获取请求对象类型并创建实例
		reqType := methodType.In(2)
		req := reflect.New(reqType.Elem())

		// 根据HTTP方法解析请求
		if methodInfo.HTTPMethod == http.MethodGet || methodInfo.HTTPMethod == http.MethodDelete {
			// 从Query参数解析
			if err := parseQueryParams(r, req.Interface()); err != nil {
				slog.Error("parse query error", "method", methodInfo.MethodName, "err", err)
				response.JSON(w, response.Resp{
					Code: ecode.ErrCodeParam,
					Msg:  err.Error(),
				})
				return
			}
		} else {
			// 从JSON Body解析
			if err := json.NewDecoder(r.Body).Decode(req.Interface()); err != nil {
				slog.Error("decode body error", "method", methodInfo.MethodName, "err", err)
				response.JSON(w, response.Resp{
					Code: ecode.ErrCodeParam,
					Msg:  err.Error(),
				})
				return
			}
		}

		// 验证请求对象
		if config != nil && config.Validator != nil {
			if err := config.Validator.StructCtx(r.Context(), req.Interface()); err != nil {
				slog.Error("validation error", "method", methodInfo.MethodName, "err", err)
				response.JSON(w, response.Resp{
					Code: ecode.ErrCodeParam,
					Msg:  err.Error(),
				})
				return
			}
		}

		args = append(args, req)

		// 调用方法
		results := method.Func.Call(args)

		// 处理返回值
		handleMethodResults(w, results, methodInfo.MethodName)
	}
}

// parseQueryParams 从URL Query参数解析到结构体
func parseQueryParams(r *http.Request, v interface{}) error {
	// 将Query参数转换为JSON
	data := make(map[string]interface{})
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			data[key] = values[0]
		}
	}

	// 使用JSON Marshal/Unmarshal来处理类型转换
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, v)
}

// handleMethodResults 处理服务方法的返回值
// 期望的返回值: (*Data, error)
func handleMethodResults(w http.ResponseWriter, results []reflect.Value, methodName string) {
	if len(results) < 2 {
		slog.Error("invalid return signature", "method", methodName, "numResults", len(results))
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  "internal server error",
		})
		return
	}

	// 获取返回的error (最后一个返回值)
	var err error
	if !results[1].IsNil() {
		err = results[1].Interface().(error)
	}

	// 处理错误
	if err != nil {
		slog.Error("method error", "method", methodName, "err", err)
		response.JSON(w, response.Resp{
			Code: ecode.ErrCodeServer,
			Msg:  err.Error(),
		})
		return
	}

	// 返回数据
	data := results[0].Interface()
	response.JSON(w, response.Resp{
		Code: ecode.OK,
		Msg:  "success",
		Data: data,
	})
}

// RegisterServiceRoutes 自动注册服务的所有公开方法作为HTTP路由
func RegisterServiceRoutes(group *Group, service interface{}, pathPrefix string, config *HandlerConfig) {
	methods := DiscoverServiceMethods(service)

	for _, methodInfo := range methods {
		// 构建完整路由路径
		fullPath := pathPrefix + methodInfo.Path

		// 创建处理程序
		handler := CreateAutoHandler(methodInfo, service, config)

		// 注册路由
		pattern := methodInfo.HTTPMethod + " " + fullPath
		slog.Info("auto register route", "pattern", pattern, "method", methodInfo.MethodName)
		group.HandleFunc(pattern, handler)
	}
}
