package hs

/*
自动路由注册功能使用指南

1. Service 方法命名规则:
   - 方法名以 HTTP 方法前缀开头 (Get, Post, Put, Delete, Patch, Create, Update, List)
   - 示例:
     - GetUser      -> GET  /user
     - GetUserList  -> GET  /user/list
     - CreateUser   -> POST /user
     - UpdateUser   -> PUT  /user
     - DeleteUser   -> DELETE /user
     - ListUsers    -> GET  /users

2. Service 方法签名规则:
   所有方法应遵循这个签名模式:

   func (s *YourService) MethodName(ctx context.Context, req *RequestDTO) (*ResponseDTO, error) {
       // 处理业务逻辑
   }

   或对于列表方法:

   func (s *YourService) ListData(ctx context.Context, req *ListRequest) ([]*DataDTO, int64, error) {
       // 返回: 数据列表, 总数, 错误
   }

   或仅返回错误:

   func (s *YourService) DeleteUser(ctx context.Context, req *DeleteRequest) error {
       // 返回: 错误
   }

3. 在 router.go 中注册服务路由:

   func RegisterRoutes(container *Container) http.Handler {
       api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)

       // 自动注册 I18n 服务的所有方法作为路由
       i18nConfig := &hs.HandlerConfig{
           Validator: container.svc.Validator,  // 可选的验证器
       }
       hs.RegisterServiceRoutes(api, container.svc.I18n, "/i18n", i18nConfig)

       // 或不需要验证器:
       // hs.RegisterServiceRoutes(api, container.svc.I18n, "/i18n", nil)

       return handlerChain
   }

4. 自动生成的路由示例:
   假设 I18nService 有以下方法:
   - Get(ctx, req)      -> GET    /api/i18n/get
   - List(ctx, req)     -> GET    /api/i18n/list
   - Create(ctx, req)   -> POST   /api/i18n/create
   - Update(ctx, req)   -> PUT    /api/i18n/update
   - Delete(ctx, req)   -> DELETE /api/i18n/delete

5. 请求绑定规则:
   - GET/DELETE 方法: 参数从 URL Query 字符串绑定
     例如: GET /api/user/get?id=123

   - POST/PUT/PATCH 方法: 参数从 JSON Body 绑定
     例如: POST /api/user/create
           Content-Type: application/json
           { "name": "John", "email": "john@example.com" }

6. 响应格式:
   所有响应都会自动包装成标准格式:

   {
     "code": 0,                    // 0 表示成功, 其他为错误码
     "msg": "success",             // 消息
     "data": { ... }               // 返回的数据
   }

   列表响应:
   {
     "code": 0,
     "msg": "success",
     "data": {
       "list": [ ... ],            // 数据数组
       "total": 100                // 总数
     }
   }

7. 验证:
   如果传入了 HandlerConfig 包含 Validator, 会自动验证请求 DTO:

   type GetUserReq struct {
       ID int64 `json:"id" validate:"required"`
   }

   如果验证失败,会返回 ErrCodeParam 错误

8. 优势:
   - 减少重复的 Handler 代码
   - 自动化路由注册,无需手动指定每个路由
   - 保持一致的请求/响应格式
   - 支持自动参数验证
   - 驼峰命名自动转换为斜杠路径
*/
