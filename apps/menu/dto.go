package menu

type MenuListReq struct{}

type MenuListResp struct {
	List []*Menu `json:"list"`
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
	Sort      int            `json:"sort"`      // 排序
	RoleIds   string         `json:"role_ids"`  // 角色ID，多个用逗号隔开
	Children  []*MenuItem    `json:"children"`
}

type MenuCreateReq struct {
	ParentID  int64  `json:"parent_id"`
	Path      string `json:"path" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Component string `json:"component" validate:"required"`
	Redirect  string `json:"redirect"`
	Meta      string `json:"meta"`
	Icon      string `json:"icon"`
	Sort      int    `json:"sort"`
	Operator  string `json:"operator"`
}

type MenuUpdateReq struct {
	ID        int64  `json:"id" validate:"required"`
	ParentID  int64  `json:"parent_id"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Meta      string `json:"meta"`
	Icon      string `json:"icon"`
	Sort      int    `json:"sort"`
	Operator  string `json:"operator"`
}

type DeleteMenuReq struct {
	IDs string `json:"ids" validate:"required"`
}

type MenuTreeResp struct {
	List []*MenuItem `json:"list"`
}
