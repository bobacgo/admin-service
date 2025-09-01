package model

type Role struct {
	Model
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      int8   `json:"status"`
}

func (Role) TableName() string {
	return "roles"
}
