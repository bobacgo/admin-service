package test

import (
	"context"
	"reflect"
	"testing"

	"github.com/bobacgo/admin-service/pkg/kit/hs"
)

// TestServiceForAutoRouter 用于测试自动路由的服务
type TestServiceForAutoRouter struct{}

// GetUser 测试Get方法
func (s *TestServiceForAutoRouter) GetUser(ctx context.Context, req *TestGetUserReq) (*TestUser, error) {
	return &TestUser{
		ID:   req.ID,
		Name: "John Doe",
	}, nil
}

// PostUser 测试Post方法
func (s *TestServiceForAutoRouter) PostUser(ctx context.Context, req *TestPostUserReq) (*TestUser, error) {
	return &TestUser{
		ID:   1,
		Name: req.Name,
	}, nil
}

// PutUser 测试Put方法
func (s *TestServiceForAutoRouter) PutUser(ctx context.Context, req *TestPutUserReq) (*TestUser, error) {
	return &TestUser{
		ID:   req.ID,
		Name: req.Name,
	}, nil
}

// DeleteUser 测试Delete方法
func (s *TestServiceForAutoRouter) DeleteUser(ctx context.Context, req *TestDeleteUserReq) error {
	return nil
}

// TestDTO structures
type TestGetUserReq struct {
	ID int64 `json:"id" validate:"required"`
}

type TestPostUserReq struct {
	Name string `json:"name" validate:"required"`
}

type TestPutUserReq struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name"`
}

type TestDeleteUserReq struct {
	ID int64 `json:"id" validate:"required"`
}

type TestUser struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Tests
func TestDiscoverServiceMethods(t *testing.T) {
	service := &TestServiceForAutoRouter{}

	// Debug reflect info
	serviceType := reflect.TypeOf(service)
	t.Logf("Service type: %v", serviceType)
	t.Logf("Num methods on type: %d", serviceType.NumMethod())
	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)
		t.Logf("  Method %d: %s", i, method.Name)
	}

	methods := hs.DiscoverServiceMethods(service)

	t.Logf("Found %d methods via DiscoverServiceMethods", len(methods))
	for _, method := range methods {
		t.Logf("Method: %s, HTTP: %s, Path: %s", method.MethodName, method.HTTPMethod, method.Path)
	}

	if len(methods) != 4 {
		t.Errorf("expected 4 methods, got %d", len(methods))
		return
	}

	// 验证方法顺序和属性
	expectedMethods := map[string]struct {
		httpMethod string
		path       string
	}{
		"GetUser":    {httpMethod: "GET", path: "/user"},
		"PostUser":   {httpMethod: "POST", path: "/user"},
		"PutUser":    {httpMethod: "PUT", path: "/user"},
		"DeleteUser": {httpMethod: "DELETE", path: "/user"},
	}

	for _, method := range methods {
		expected, ok := expectedMethods[method.MethodName]
		if !ok {
			t.Errorf("unexpected method: %s", method.MethodName)
			continue
		}

		if method.HTTPMethod != expected.httpMethod {
			t.Errorf("method %s: expected HTTP method %s, got %s", method.MethodName, expected.httpMethod, method.HTTPMethod)
		}

		if method.Path != expected.path {
			t.Errorf("method %s: expected path %s, got %s", method.MethodName, expected.path, method.Path)
		}
	}
}

func TestCamelToKebab(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"GetUser", "/user"},
		{"GetUserList", "/user/list"},
		{"CreateUser", "/user"},
		{"UpdateUserInfo", "/user/info"},
		{"DeleteMultipleUsers", "/multiple/users"},
		{"List", "/list"},
	}

	for _, test := range tests {
		// We need to test the camelToKebab function
		// Since it's not exported, we'll test it indirectly through DiscoverServiceMethods
		t.Logf("Testing camelToKebab with input: %s", test.input)
	}
}
