package dto

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserListReq struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type GetUserReq struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}