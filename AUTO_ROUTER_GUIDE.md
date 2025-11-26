# Auto-Router Implementation - Multi-Service Merge Complete

## Overview
所有四个服务（User, Menu, Role, I18n）现已完全使用自动路由模式，handler 层已被消除。Service 方法直接注册为 HTTP 路由。

## Service Method Naming Convention

### Standard HTTP Methods (Auto-Routed)
- `GetXxx(ctx, req) (resp, error)` → `GET /xxx`
- `PostXxx(ctx, req) (resp, error)` → `POST /xxx`
- `PutXxx(ctx, req) (resp, error)` → `PUT /xxx`
- `DeleteXxx(ctx, req) (resp, error)` → `DELETE /xxx`
- `PatchXxx(ctx, req) (resp, error)` → `PATCH /xxx`

### Current Implementations

#### User Service (`apps/user/service.go`)
- `GetOne(ctx, *GetUserReq) (*User, error)` → `GET /api/user/one`
- `GetList(ctx, *UserListReq) (*PageResp[User], error)` → `GET /api/user/list`
- `PostCreate(ctx, *User) (*User, error)` → `POST /api/user/create`
- `PutUpdate(ctx, *User) (*User, error)` → `PUT /api/user/update`
- `DeleteDel(ctx, *DeleteUserReq) (interface{}, error)` → `DELETE /api/user/del`
- `Login(ctx, *LoginReq) (map[string]string, error)` → Custom handler → `POST /api/login`
- `Logout(ctx) error` → Custom handler → `POST /api/logout`

#### Menu Service (`apps/menu/service.go`)
- `GetOne(ctx, *GetMenuReq) (*Menu, error)` → `GET /api/menu/one`
- `GetList(ctx, *MenuListReq) (*PageResp[Menu], error)` → `GET /api/menu/list`
- `PostCreate(ctx, *MenuCreateReq) (*Menu, error)` → `POST /api/menu/create`
- `PutUpdate(ctx, *MenuUpdateReq) (*Menu, error)` → `PUT /api/menu/update`
- `DeleteDel(ctx, *DeleteMenuReq) (interface{}, error)` → `DELETE /api/menu/del`
- `GetTree(ctx, interface{}) ([]*MenuItem, error)` → Custom handler → `GET /api/get-menu-list-i18n`

#### Role Service (`apps/role/service.go`)
- `GetOne(ctx, *GetRoleReq) (*Role, error)` → `GET /api/role/one`
- `GetList(ctx, *RoleListReq) (*PageResp[Role], error)` → `GET /api/role/list`
- `PostCreate(ctx, *RoleCreateReq) (*Role, error)` → `POST /api/role/create`
- `PutUpdate(ctx, *RoleUpdateReq) (*Role, error)` → `PUT /api/role/update`
- `DeleteDel(ctx, *DeleteRoleReq) (interface{}, error)` → `DELETE /api/role/del`

#### I18n Service (`apps/i18n/service.go`)
- `GetOne(ctx, *GetI18nReq) (*I18n, error)` → `GET /api/i18n/one`
- `GetList(ctx, *I18nListReq) (*PageResp[I18n], error)` → `GET /api/i18n/list`
- `PostCreate(ctx, *I18nCreateReq) (*I18n, error)` → `POST /api/i18n/create`
- `PutUpdate(ctx, *I18nUpdateReq) (*I18n, error)` → `PUT /api/i18n/update`
- `DeleteDel(ctx, *DeleteI18nReq) (interface{}, error)` → `DELETE /api/i18n/del`

## DTO Pattern

All DTOs follow this pattern with validation tags:
```go
type DeleteXxxReq struct {
	IDs string `json:"ids" query:"ids" validate:"required"`
}
```

For complex operations:
```go
type PostCreateReq struct {
	Field1 string `json:"field1" validate:"required"`
	Field2 int    `json:"field2" validate:"gt=0"`
	// ... other fields
}
```

## Response Format
All auto-routed responses are wrapped in standard format:
```go
{
	"code": 0,           // 0 for success, error code otherwise
	"msg": "ok",         // Error message
	"data": {            // Actual response data
		// ... service method return value
	}
}
```

## Implementation Details

### Auto-Router Core (`pkg/kit/hs/auto_router.go`)
- Discovers public methods with HTTP method prefixes (Get, Post, Put, Delete, Patch)
- Extracts HTTP method from prefix and generates path from method name
- Creates handlers that automatically bind request parameters to DTOs
- Wraps responses in standard format with automatic error handling

### Request Parameter Binding
- **GET/DELETE**: Parameters from URL query string (requires `query` tag in DTO)
- **POST/PUT/PATCH**: Parameters from JSON request body (requires `json` tag in DTO)

### Validation
- All DTOs are automatically validated using validator/v10
- Validation is performed before the service method is called
- Validation errors return error code 100001 (ErrCodeParam)

### Special Handlers
For methods that don't follow the HTTP prefix pattern (like Login, Logout, GetTree):
- Create custom handler functions in `router.go`
- Register them using `api.HandleFunc()` directly
- Use same response wrapper for consistency

## Adding New Auto-Routed Methods

1. **Rename method** to match HTTP prefix pattern (GetXxx, PostXxx, etc.)
2. **Create DTO** in `apps/xxx/dto.go` with validation tags
3. **Add validator field** to service struct
4. **Implement method** with signature: `func(ctx context.Context, req *ReqDTO) (*RespDTO, error)`
5. **Update IOC** in `apps/ioc.go` to pass validator to service constructor
6. **Remove manual routes** from `apps/router.go` - auto-router will handle it

## Files Modified

- `apps/user/service.go` - Added new method signatures, validator, Login/Logout
- `apps/user/dto.go` - Added DeleteUserReq
- `apps/user/handler.go` - Deprecated
- `apps/menu/service.go` - Added new method signatures, validator, GetTree
- `apps/menu/dto.go` - Added DeleteMenuReq  
- `apps/menu/handler.go` - Deprecated
- `apps/role/service.go` - Added new method signatures, validator
- `apps/role/dto.go` - Added DeleteRoleReq
- `apps/role/handler.go` - Deprecated
- `apps/i18n/service.go` - Added new method signatures, validator, GetList
- `apps/i18n/dto.go` - Added DeleteI18nReq
- `apps/ioc.go` - Updated service initialization with validator
- `apps/api/api.go` - Removed handler fields
- `apps/router.go` - Updated to use RegisterServiceRoutes and custom handlers

## Testing
- Auto-router method discovery: ✅ PASS
- Camel-to-kebab path conversion: ✅ PASS
- Full project build: ✅ SUCCESS
