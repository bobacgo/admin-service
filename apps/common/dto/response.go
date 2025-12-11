package dto

type PageResp[T any] struct {
	Total int64 `json:"total"`
	List  []*T  `json:"list"`
}

func NewPageResp[T any](total int64, list []*T) *PageResp[T] {
	if list == nil {
		list = make([]*T, 0)
	}
	return &PageResp[T]{
		Total: total,
		List:  list,
	}
}
