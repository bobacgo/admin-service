# 合并 Handler 到 Service 计划

## 目标

消除 i18n handler，利用自动路由功能直接从 service 方法生成 HTTP 路由。

## 当前状态

- 有独立的 `I18nHandler` 处理 HTTP 请求
- 有 `I18nService` 处理业务逻辑
- 两层需要参数解析、验证和响应封装

## 改动方案

### 1. 更新 `apps/i18n/service.go`

将服务方法改为符合自动路由的签名 `(ctx context.Context, req *ReqType) (*RespType, error)`:

```go
package i18n

import (
	"context"
	"time"

	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/go-playground/validator/v10"
)

type I18nService struct {
	repo      *I18nRepo
	validator *validator.Validate
}

func NewI18nService(r *I18nRepo, v *validator.Validate) *I18nService {
	return &I18nService{repo: r, validator: v}
}

// Get 获取单个i18n记录
func (s *I18nService) GetOne(ctx context.Context, req *GetI18nReq) (*I18n, error) {
	return s.repo.FindOne(ctx, req)
}

// Post 创建i18n记录
func (s *I18nService) PostCreate(ctx context.Context, req *I18nCreateReq) (*I18n, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	i18n := &I18n{
		Class: req.Class,
		Lang:  req.Lang,
		Key:   req.Key,
		Value: req.Value,
		Model: model.Model{
			CreatedAt: time.Now().Unix(),
		},
	}

	if err := s.repo.Create(ctx, i18n); err != nil {
		return nil, err
	}

	return i18n, nil
}

// Put 更新i18n记录
func (s *I18nService) PutUpdate(ctx context.Context, req *I18nUpdateReq) (*I18n, error) {
	if err := s.validator.StructCtx(ctx, req); err != nil {
		return nil, err
	}

	i18n := &I18n{
		Class: req.Class,
		Lang:  req.Lang,
		Value: req.Value,
		Model: model.Model{
			ID: req.ID,
		},
	}

	if err := s.repo.Update(ctx, i18n); err != nil {
		return nil, err
	}

	return i18n, nil
}

// Delete 删除i18n记录
func (s *I18nService) DeleteBatch(ctx context.Context, req *DeleteI18nReq) (interface{}, error) {
	return nil, s.repo.Delete(ctx, req.IDs)
}
```

### 2. 添加 DTO 到 `apps/i18n/dto.go`

```go
type DeleteI18nReq struct {
	IDs string `json:"ids" query:"ids" validate:"required"`
}
```

### 3. 更新 `apps/ioc.go`

修改 I18nService 的初始化，传递 validator：

```go
// 从这样:
i18nSvc := i18n.NewI18nService(i18nRepo)

// 改为:
i18nSvc := i18n.NewI18nService(i18nRepo, service.GetValidator())
```

### 4. 保留 `apps/router.go`

已经配置好自动路由注册：

```go
i18nConfig := &hs.HandlerConfig{
    Validator: container.svc.Validator,
}
hs.RegisterServiceRoutes(api, container.svc.I18n, "/i18n", i18nConfig)
```

### 5. 删除 `apps/i18n/handler.go`

由于 service 现在直接处理 HTTP 请求，handler 文件已不需要。

### 6. 删除 `apps/api/api.go` 中的 I18nHandler

```go
// 删除这行:
I18n *i18n.I18nHandler

// 删除这行:
I18n: i18n.NewI18nHandler(svc.I18n, svc.Validator),
```

## 路由映射

自动生成的路由：

- `GET /api/i18n/get` → `I18nService.Get()`
- `POST /api/i18n/post` → `I18nService.Post()`
- `PUT /api/i18n/put` → `I18nService.Put()`
- `DELETE /api/i18n/delete` → `I18nService.Delete()`

## 注意事项

1. **参数验证移到 service** - 验证逻辑从 handler 移到 service 方法
2. **删除请求DTO** - Delete 需要新增 `DeleteI18nReq` DTO
3. **统一返回值** - 所有 service 方法返回 `(*Data, error)`
4. **自动路由处理参数绑定** - 不再需要手动处理 JSON/Query 解析

## 优势

- ✅ 减少代码重复 (消除 handler 层)
- ✅ 统一的参数绑定和验证
- ✅ 自动化路由注册
- ✅ 简化架构，service 直接对应 HTTP 端点
