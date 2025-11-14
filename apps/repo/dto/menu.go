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

type MenuItem struct {
	ID        int64          `json:"id"`        // ID
	ParentID  int64          `json:"parent_id"` // 父ID
	Path      string         `json:"path"`      // 路径
	Name      string         `json:"name"`      // 名称
	Component string         `json:"component"` // 组件
	Redirect  string         `json:"redirect"`  // 重定向
	Meta      map[string]any `json:"meta"`      // 元数据
	Icon      string         `json:"icon"`      // 图标
	Sort      int64          `json:"sort"`      // 排序
	Children  []*MenuItem    `json:"children"`
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
