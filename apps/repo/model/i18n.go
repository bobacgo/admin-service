package model

import "github.com/bobacgo/orm"

const I18nTable = "i18n"

type I18n struct {
	Model
	Class string `json:"class"` // 分类
	Lang  string `json:"lang"`  // 语言
	Key   string `json:"key"`   // 键
	Value string `json:"value"` // 值
}

const (
	Class string = "class"
	Lang  string = "lang"
	Key   string = "key"
	Value string = "value"
)

func (m *I18n) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: Id, Result: &m.ID, Value: m.ID},
		{Column: Class, Result: &m.Class, Value: m.Class},
		{Column: Lang, Result: &m.Lang, Value: m.Lang},
		{Column: Key, Result: &m.Key, Value: m.Key},
		{Column: Value, Result: &m.Value, Value: m.Value},
		{Column: CreatedAt, Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: UpdatedAt, Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}