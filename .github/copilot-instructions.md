# Admin Service - AI 编码助手指南

## 架构概览

这是一个基于 **Go** 的后台管理服务，搭配 Vue 3 前端，核心特性是采用了自定义的**自动路由框架**，消除了传统的 handler 层。后端遵循清晰的架构模式：服务层、仓储层和数据访问层。

### 核心组件
- **后端**: 使用自定义 `pkg/kit/hs` 框架的 Go HTTP 服务器（端口 8080）
- **前端**: Vue 3 + Vite + TDesign UI（位于 `web/` 目录）
- **数据库**: MySQL，使用基于时间戳的字段（BIGINT 存储 Unix 时间戳）
- **ORM**: 自定义 ORM (`github.com/bobacgo/orm`)，提供流式 API

## 自动路由系统（核心！）

**服务方法直接暴露为 HTTP 端点**，无需 handler 层。方法命名决定 HTTP 路由：

### 方法命名约定
```go
// 模式: {HTTP方法}{路由路径}
func (s *Service) GetOne(ctx context.Context, req *GetReq) (*Model, error)        // → GET /api/{service}/one
func (s *Service) GetList(ctx context.Context, req *ListReq) (*PageResp, error)  // → GET /api/{service}/list
func (s *Service) Post(ctx context.Context, req *Model) (*Model, error)          // → POST /api/{service}
func (s *Service) Put(ctx context.Context, req *Model) (*Model, error)           // → PUT /api/{service}
func (s *Service) Delete(ctx context.Context, req *DeleteReq) (any, error)       // → DELETE /api/{service}
```

**方法名转换规则**: `GetUserInfo` → `/user/info`，`PostCreate` → `/create`

### 方法签名要求
- **必须 3 个输入参数**: (接收者, `context.Context`, 请求 DTO 指针)
- **必须 2 个输出参数**: (响应指针或 `any`, `error`)
- GET/DELETE: 从 query 参数绑定（需使用 `query:"field"` 标签）
- POST/PUT/PATCH: 从 JSON body 绑定（需使用 `json:"field"` 标签）

### 注册模式（在 `apps/router.go` 中）
```go
api := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)
hs.RegisterService(api, "/user", container.svc.User)  // 自动发现并注册服务方法
```

### 自定义 Handler（非标准方法）
对于 `Login()` 或 `GetTree()` 等不遵循自动路由的方法，需创建手动 handler：
```go
api.HandleFunc("POST /login", makeLoginHandler(container.svc.User))
```

详细实现请参考 `AUTO_ROUTER_GUIDE.md`。

## 项目结构

```
apps/                    # 业务模块（user, menu, role, i18n）
  {module}/
    dto.go              # 请求/响应 DTO，带验证标签
    model.go            # 数据库模型，必须实现 ORM Mapping() 方法
    repo.go             # 仓储层（CRUD 操作）
    service.go          # 服务层（自动路由的方法）
  ioc.go                # 依赖注入容器
  router.go             # 路由注册
  ecode/                # 集中的错误码定义
  repo/
    data/data.go        # 数据库连接（读取 DB_* 环境变量）
    dto/                # 共享 DTO（PageReq, PageResp）
    model/              # 共享 model 结构（Model 基础类型）
cmd/server/main.go      # 应用程序入口
pkg/kit/hs/             # 自定义 HTTP 框架
  auto_router.go        # 自动路由发现/注册
  middleware.go         # Logger、Cors、Auth 中间件
  response/r.go         # 标准 JSON 响应包装器
web/                    # Vue 3 前端（独立的 npm 项目）
```

## 开发工作流

### 运行服务
```bash
# 后端（在项目根目录）
go run cmd/server/main.go

# 前端（在 web/ 目录）
cd web && npm run dev
```

### 数据库设置
```bash
# 环境变量（显示默认值）
export DB_USER=root
export DB_PASS=
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_NAME=admin_db

# 初始化数据库
mysql -u root < db.sql
```

### 测试
测试使用 `net/http/httptest`，包含完整的服务设置：
```bash
go test ./test/...
```

每个测试文件（`test/{module}_test.go`）都会创建完整的服务栈和数据库连接。

## 代码模式与约定

### 带验证的 DTO
```go
type CreateReq struct {
    Field1 string `json:"field1" validate:"required"`
    Field2 int    `json:"field2" validate:"gt=0"`
}

type DeleteReq struct {
    IDs string `json:"ids" query:"ids" validate:"required"`  // 批量删除使用逗号分隔的 ID
}
```

### Model Mapping（ORM 必需）
每个 model **必须**实现 `Mapping()` 方法，将结构体字段绑定到数据库列：
```go
func (m *User) Mapping() []*orm.Mapping {
    return []*orm.Mapping{
        {Column: "id", Result: &m.ID, Value: m.ID},
        {Column: "account", Result: &m.Account, Value: m.Account},
        // ... 所有字段
    }
}
```

### 使用自定义 ORM 的仓储模式
```go
// 创建
id, err := INSERT(row).INTO("users").Omit("id").Exec(ctx, db)

// 查询
err := SELECT1(row).FROM("users").WHERE(map[string]any{"id": 1}).Query(ctx, db)

// 分页列表
rows, total, err := SELECT(new(User)).FROM("users").
    WHERE(where).
    LIMIT(req.PageSize).
    OFFSET((req.Page - 1) * req.PageSize).
    QueryPage(ctx, db)

// 更新
err := UPDATE(row).SET("account", "password").WHERE(map[string]any{"id": row.ID}).Exec(ctx, db)

// 删除
err := DELETE().FROM("users").WHERE(map[string]any{"id IN": ids}).Exec(ctx, db)
```

### 标准响应格式
所有自动路由的响应都会被包装：
```json
{
  "code": 0,
  "msg": "ok",
  "data": { /* 实际响应数据 */ }
}
```

错误码定义在 `apps/ecode/err_code.go` 中：
- `0`: 成功
- `100001`: 参数验证错误
- `100100`: 登录失败（用户名/密码）

### 中间件应用
```go
group := hs.NewGroup("/api", mux, hs.Logger, hs.Cors)  // 应用中间件
```
标准中间件：`Logger`（panic 恢复 + 计时）、`Cors`（完整的 CORS 头）

## 添加新功能

### 创建新的服务模块
1. **创建目录**: `apps/{module}/`
2. **定义 model**，实现 `Mapping()` 方法（`model.go`）
3. **创建 DTO**，添加验证标签（`dto.go`）
4. **实现仓储**，使用 ORM 流式 API（`repo.go`）
5. **编写服务**，使用自动路由的方法命名（`service.go`）
6. **依赖注入**，在 `apps/ioc.go` 中连接（初始化 repo → service → 添加到容器）
7. **注册路由**，在 `apps/router.go` 中：`hs.RegisterService(api, "/{module}", svc)`
8. **添加表**，更新 `db.sql`

### 数据库 Schema 要求
- 主键：`id BIGINT AUTO_INCREMENT`
- 时间戳：`created_at`、`updated_at` BIGINT（Unix 时间戳）
- 状态字段使用 `TINYINT NOT NULL DEFAULT 1`
- 所有字段必须添加 `COMMENT` 中文注释
- 字符串字段使用合理的长度（避免统一 VARCHAR(255)）
- 索引命名规范：`idx_{table}_{column}` 或 `uq_{table}_{column}`
- 关联表使用有意义的字段名（如 `user_id`, `role_id`，而非 `r1_id`, `r2_id`）
- 根据需要创建唯一/普通/复合索引

## 重要提示

- **无 handler 层**：服务方法通过反射直接路由
- **自动验证**：所有 DTO 在调用服务方法前自动验证
- **上下文传递**：始终通过仓储调用传递 `ctx`
- **时间处理**：使用 Unix 时间戳（`time.Now().Unix()`）而非 `time.Time`
- **前端集成**：Vue 应用期望后端在 `localhost:8080`，已启用 CORS
- **自定义 ORM**：使用基于反射的映射，需要实现 `Mapping()` 方法
- **错误处理**：从服务返回 error；自动路由器会用错误码包装它们

## 常见陷阱

1. **错误的方法签名**：自动路由要求精确的签名（3 入，2 出）
2. **缺少标签**：GET/DELETE 的 DTO 需要同时有 `json` 和 `query` 标签
3. **忘记 Mapping()**：没有 `Mapping()` 的 model 会在 ORM 操作时失败
4. **直接注册 mux**：对服务使用 `hs.RegisterService()` 而不是 `mux.HandleFunc()`
5. **时间戳类型**：MySQL schema 中使用 `BIGINT` 而非 `DATETIME`
