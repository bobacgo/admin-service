package model

import "github.com/bobacgo/orm"

// common columns
const (
	Id        string = "id"
	CreatedAt string = "created_at"
	UpdatedAt string = "updated_at"

	Status string = "status" // 状态 1:正常 2:禁用
)

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

const (
	R1Id string = "r1_id"
	R2Id string = "r2_id"
)

func (m *Relation) Mappping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: Id, Result: &m.ID, Value: m.ID},
		{Column: R1Id, Result: &m.R1Id, Value: m.R1Id},
		{Column: R2Id, Result: &m.R2Id, Value: m.R2Id},
		{Column: CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}