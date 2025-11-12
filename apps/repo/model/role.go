package model

import "github.com/bobacgo/orm"

const (
	RoleCodeAdmin string = "admin"
	RoleTable     string = "roles"
)

type Role struct {
	Model
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
}

const (
	Code        string = "code"
	Description string = "description"
)

func (m *Role) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: Id, Result: &m.ID, Value: m.ID},
		{Column: Code, Result: &m.Code, Value: m.Code},
		{Column: Description, Result: &m.Description, Value: m.Description},
		{Column: Status, Result: &m.Status, Value: m.Status},
		{Column: CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}
