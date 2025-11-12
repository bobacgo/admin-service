package dto

type GetMenuReq struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
}

type MenuListReq struct {
	PageReq
	Path string `json:"path"`
	Name string `json:"name"`
}

type MenuCreateReq struct {
	Path      string `json:"path" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Meta      string `json:"meta"`
	Icon      string `json:"icon"`
}

type MenuUpdateReq struct {
	ID        int64  `json:"id" validate:"required"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Meta      string `json:"meta"`
	Icon      string `json:"icon"`
}
