package model

import "github.com/bobacgo/orm"

const I18nTable = "i18n"

type I18n struct {
	Model
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (m *I18n) Mapping() []*orm.Mapping {
	return []*orm.Mapping{
		{Column: "id", Result: &m.ID, Value: m.ID},
		{Column: "class", Result: &m.Class, Value: m.Class},
		{Column: "lang", Result: &m.Lang, Value: m.Lang},
		{Column: "key", Result: &m.Key, Value: m.Key},
		{Column: "value", Result: &m.Value, Value: m.Value},
		{Column: "created_at", Result: &m.CreatedAt, Value: m.CreatedAt},
		{Column: "updated_at", Result: &m.UpdatedAt, Value: m.UpdatedAt},
	}
}