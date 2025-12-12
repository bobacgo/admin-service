package dto

import (
	"github.com/bobacgo/admin-service/apps/common/dto"
)

type LoginReq struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserListReq struct {
	dto.PageReq
	Account string `json:"account"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Status  int8   `json:"status"`   // 1:正常 2:禁用
	RoleIds string `json:"role_ids"` // 角色ID，逗号分隔
}

type GetUserReq struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type DeleteUserReq struct {
	IDs string `json:"ids" query:"ids" validate:"required"`
}

// 非关联事件字段更新
// 【更新】用户相关
type UpdateUserReq struct {
	dto.AdminUpdate
	Phone string `json:"phone"`                            // 手机号
	Email string `json:"email" validate:"omitempty,email"` // 邮箱
}

// 有关联事件的字段更新
// 【更新】用户登录信息
type UpdateLoginInfoReq struct {
	Id      int64  // 用户ID
	LoginAt int64  // 登录时间
	LoginIp string // 登录IP
}

// 有关联事件的字段更新
// 【更新】用户状态
type UpdateUserStatusReq struct {
	dto.AdminUpdate
	Status int8 `json:"status" validate:"required,oneof=1 2"` // 状态 1:正常 2:禁用
}

// 有关联事件的字段更新
// 【更新】用户角色
type UpdateUserRoleReq struct {
	dto.AdminUpdate
	RoleIds string `json:"role_ids" validate:"required"` // 角色ID，逗号分隔
}

// 有关联事件的字段更新
// 【更新】用户密码
type UpdateUserPasswordReq struct {
	dto.AdminUpdate
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}