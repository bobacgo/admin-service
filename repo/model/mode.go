package model

type Model struct {
	ID        uint32 `json:"id"`
	CreatedAt int64  `json:"created_at"` // 时间戳
	UpdatedAt int64  `json:"updated_at"` // 时间戳
}