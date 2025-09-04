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
	Lang  string `json:"lang" validate:"required"`
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type I18nUpdateReq struct {
	ID    int64  `json:"id" validate:"required"`
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Value string `json:"value"`
}
