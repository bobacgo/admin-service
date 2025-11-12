package dto

type GetRoleReq struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

type RoleListReq struct {
	PageReq
	Code   string `json:"code"`
	Status string `json:"status"` // 逗号分隔
}

type RoleCreateReq struct {
	Code        string `json:"code" validate:"required"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
}

type RoleUpdateReq struct {
	ID          int64  `json:"id" validate:"required"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
}
