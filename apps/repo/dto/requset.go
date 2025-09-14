package dto

import (
	"strconv"
	"strings"
)

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

func StringToIds(s string) []int64 {
	ids := make([]int64, 0)
	if s == "" {
		return ids
	}
	for _, v := range strings.Split(s, ",") {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

type PageReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (p *PageReq) Limit() (int, int) {
	// 任意为负数不分页
	if p.Page <= 0 {
		return 0, 0
	}
	if p.PageSize <= 0 {
		return 0, 0
	}

	// offset limit
	return (p.Page - 1) * p.PageSize, p.PageSize
}
