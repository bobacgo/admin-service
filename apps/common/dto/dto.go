package dto

type AdminUpdate struct {
	Id        int64  `json:"id" validate:"required"`
	Operator  string `json:"-"` // 操作人，  服务端设置
	UpdatedAt int64  `json:"-"` // 更新时间，服务端设置
}
