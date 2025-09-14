package model

type MenuMeta struct {
	Title map[string]string `json:"title"`
}

type Menu struct {
	Model
	Path      string  `json:"path"`
	Name      string  `json:"name"`
	Component string  `json:"component"`
	Redirect  string  `json:"redirect"`
	Meta      string  `json:"meta"`
	Icon      string  `json:"icon"`
	Children  []*Menu `json:"children"`
}

func TableName() string {
	return "menus"
}
