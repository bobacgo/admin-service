package dto

type GetI18nReq struct {
	ID    int64  `json:"id"`
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
}

type I18nListReq struct {
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
}

type I18nCreateReq struct {
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
