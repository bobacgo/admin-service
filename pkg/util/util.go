package util

import "reflect"

func IsZero(x any) bool {
	if x == nil {
		return true
	}
	return reflect.ValueOf(x).IsZero()
}
