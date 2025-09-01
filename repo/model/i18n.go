package model

type I18n struct {
	Model
	Class string `json:"class" gorm:"column:class"`
	Lang  string `json:"lang" gorm:"column:lang"`
	Key   string `json:"key" gorm:"column:key"`
	Value string `json:"value" gorm:"column:value"`
}

func (I18n) TableName() string {
	return "i18n"
}

func (I18n) Mapping(ptr bool) map[string]func(*I18n) any {
	return map[string]func(*I18n) any{
		"id": func(i *I18n) any {
			if ptr {
				return &i.ID
			}
			return i.ID
		},
		"class": func(i *I18n) any {
			if ptr {
				return &i.Class
			}
			return i.Class
		},
		"lang": func(i *I18n) any {
			if ptr {
				return &i.Lang
			}
			return i.Lang
		},
		"key": func(i *I18n) any {
			if ptr {
				return &i.Key
			}
			return i.Key
		},
		"value": func(i *I18n) any {
			if ptr {
				return &i.Value
			}
			return i.Value
		},
		"created_at": func(i *I18n) any {
			if ptr {
				return &i.CreatedAt
			}
			return i.CreatedAt
		},
		"updated_at": func(i *I18n) any {
			if ptr {
				return &i.UpdatedAt
			}
			return i.UpdatedAt
		},
	}
}
