package role

import "github.com/bobacgo/admin-service/apps/common/dto"

type GetRoleReq struct {
	ID       int64  `json:"id"`
	RoleName string `json:"role_name"`
}

type RoleListReq struct {
	dto.PageReq
	RoleName string `json:"role_name"`
	Status   string `json:"status"` // 逗号分隔
}

type RoleCreateReq struct {
	RoleName    string `json:"role_name" validate:"required"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
	Operator    string `json:"operator"`
}

type RoleUpdateReq struct {
	ID          int64  `json:"id" validate:"required"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
	Operator    string `json:"operator"`
}

type DeleteRoleReq struct {
	IDs string `json:"ids" query:"ids" validate:"required"`
}

type SaveRolePermissionsReq struct {
	RoleId  int64   `json:"role_id" validate:"required"`
	MenuIds []int64 `json:"menu_ids"`
}

type GetRolePermissionsReq struct {
	RoleId int64 `json:"role_id" query:"role_id" validate:"required"`
}

type GetRolePermissionsResp struct {
	MenuIds []int64 `json:"menu_ids"`
}
