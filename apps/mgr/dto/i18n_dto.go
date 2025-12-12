package dto

import "github.com/bobacgo/admin-service/apps/common/dto"

type GetI18nReq struct {
	ID    int64  `json:"id"`
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
}

type I18nListReq struct {
	dto.PageReq
	Class string `json:"class"`
	Lang  string `json:"lang"`
	Key   string `json:"key"`
}

type I18nCreateReq struct {
	Class    string `json:"class"`
	Lang     string `json:"lang" validate:"required"`
	Key      string `json:"key" validate:"required"`
	Value    string `json:"value" validate:"required"`
	Operator string `json:"operator"`
}

type I18nUpdateReq struct {
	ID       int64  `json:"id" validate:"required"`
	Class    string `json:"class"`
	Lang     string `json:"lang"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}

type DeleteI18nReq struct {
	IDs string `json:"ids" query:"ids" validate:"required"`
}