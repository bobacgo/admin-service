package user

import "github.com/bobacgo/admin-service/apps/repo/dto"

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserListReq struct {
	dto.PageReq
	Account string `json:"account"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Status  string `json:"status"` // 逗号分隔
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
