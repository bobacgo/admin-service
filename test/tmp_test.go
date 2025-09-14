package test

import (
	"fmt"
	"testing"
)

type Model interface {
	TableName() string
	Mapping(bool) map[string]any
}

type User struct {
	ID int
}

func (*User) TableName() string {
	return "users"
}

func (u *User) Mapping(bool) map[string]any {
	return map[string]any{
		"id": &u.ID,
	}
}

func TestName(t *testing.T) {
	one := FindOne[User]()
	fmt.Println(one)
}

func FindOne[T any, PT interface {
	*T
	Model
}]() PT {
	res := PT(new(T)) // 显式转换成 PT
	fmt.Println(res.TableName())
	mapping := res.Mapping(false)
	for _, v := range mapping {
		*v.(*int) = 12
	}
	return res
}
