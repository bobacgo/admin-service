package role

import (
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/bobacgo/orm"
)

const (
	RoleNameAdmin string = "admin"
	RoleTable     string = "roles"
)

type Role struct {
	model.Model
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
	UserCount   int64  `json:"user_count,omitempty"`
}

const (
	RoleName    string = "role_name"
	Description string = "description"
)

func (m *Role) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: model.Id, Result: &m.ID, Value: m.ID},
		{Column: RoleName, Result: &m.RoleName, Value: m.RoleName},
		{Column: Description, Result: &m.Description, Value: m.Description},
		{Column: model.Status, Result: &m.Status, Value: m.Status},
		{Column: model.Operator, Result: &m.Operator, Value: m.Operator},
		{Column: model.CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: model.UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}
