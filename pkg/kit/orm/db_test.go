package orm_test

import (
	"testing"

	"github.com/bobacgo/admin-service/pkg/kit/orm"
)

// test insert

func TestInsert(t *testing.T) {
	orm.TestInsert(t)
}

func TestInsertModel(t *testing.T) {
	orm.TestInsertModel(t)
}

func TestInsertModels(t *testing.T) {
	orm.TestInsertModels(t)
}

func TestDelete(t *testing.T) {
	orm.TestDelete(t)
}
