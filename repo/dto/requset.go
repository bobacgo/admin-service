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