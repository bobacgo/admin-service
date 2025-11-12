package model

import "github.com/bobacgo/orm"

type MenuMeta struct {
	Title map[string]string `json:"title"`
}

const MenuTable = "menus"

type Menu struct {
	Model
	Path      string  `json:"path"`      // 路径
	Name      string  `json:"name"`      // 名称
	Component string  `json:"component"` // 组件
	Redirect  string  `json:"redirect"`  // 重定向
	Meta      string  `json:"meta"`      // 元数据
	Icon      string  `json:"icon"`      // 图标
	Children  []*Menu `json:"children"`
}

const (
	Path      string = "path"
	Name      string = "name"
	Component string = "component"
	Redirect  string = "redirect"
	Meta      string = "meta"
	Icon      string = "icon"
)

func (m *Menu) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: Id, Result: &m.ID, Value: m.ID},
		{Column: Path, Result: &m.Path, Value: m.Path},
		{Column: Name, Result: &m.Name, Value: m.Name},
		{Column: Component, Result: &m.Component, Value: m.Component},
		{Column: Redirect, Result: &m.Redirect, Value: m.Redirect},
		{Column: Meta, Result: &m.Meta, Value: m.Meta},
		{Column: Icon, Result: &m.Icon, Value: m.Icon},
		{Column: CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}