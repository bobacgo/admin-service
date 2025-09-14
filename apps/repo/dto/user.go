package dto

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserListReq struct {
	PageReq
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
