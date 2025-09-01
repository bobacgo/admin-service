package model

type Model struct {
	ID        int64 `json:"id"`
	CreatedAt int64 `json:"created_at"` // 时间戳
	UpdatedAt int64 `json:"updated_at"` // 时间戳
}

// 关联关系表
type Relation struct {
	Model
	R1Id uint32 `json:"r1_id"`
	R2Id uint32 `json:"r2_id"`
}
