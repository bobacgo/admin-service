package menu

import (
	"github.com/bobacgo/admin-service/apps/repo/model"
	"github.com/bobacgo/orm"
)

type MenuMeta struct {
	Title map[string]string `json:"title"`
}

const MenuTable = "menus"

type Menu struct {
	model.Model
	ParentID  int64  `json:"parent_id"`  // 父ID
	Path      string `json:"path"`       // 路径
	Name      string `json:"name"`       // 名称
	Component string `json:"component"`  // 组件
	Redirect  string `json:"redirect"`   // 重定向
	Meta      string `json:"meta"`       // 元数据
	Icon      string `json:"icon"`       // 图标
	Sort      int    `json:"sort"`       // 排序
	RoleCodes string `json:"role_codes"` // 角色编码，多个用逗号隔开
}

const (
	ParentID  string = "parent_id"
	Path      string = "path"
	Name      string = "name"
	Component string = "component"
	Redirect  string = "redirect"
	Meta      string = "meta"
	Icon      string = "icon"
	Sort      string = "sort" // 排序
	RoleCodes string = "role_codes"
)

func (m *Menu) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: model.Id, Result: &m.ID, Value: m.ID},
		{Column: ParentID, Result: &m.ParentID, Value: m.ParentID},
		{Column: Path, Result: &m.Path, Value: m.Path},
		{Column: Name, Result: &m.Name, Value: m.Name},
		{Column: Component, Result: &m.Component, Value: m.Component},
		{Column: Redirect, Result: &m.Redirect, Value: m.Redirect},
		{Column: Meta, Result: &m.Meta, Value: m.Meta},
		{Column: Icon, Result: &m.Icon, Value: m.Icon},
		{Column: Sort, Result: &m.Sort, Value: m.Sort},
		{Column: RoleCodes, Result: &m.RoleCodes, Value: m.RoleCodes},
		{Column: model.CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: model.UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}
