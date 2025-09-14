package model

const I18nTable = "i18n"

type I18n struct {
	Model
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (*I18n) TableName() string {
	return I18nTable
}

func (m *I18n) Mapping(ptr bool) map[string]any {
	return map[string]any{
		"id":         ptrFunc(ptr, &m.ID, m.ID),
		"class":      ptrFunc(ptr, &m.Class, m.Class),
		"lang":       ptrFunc(ptr, &m.Lang, m.Lang),
		"key":        ptrFunc(ptr, &m.Key, m.Key),
		"value":      ptrFunc(ptr, &m.Value, m.Value),
		"created_at": ptrFunc(ptr, &m.CreatedAt, m.CreatedAt),
		"updated_at": ptrFunc(ptr, &m.UpdatedAt, m.UpdatedAt),
	}
}
