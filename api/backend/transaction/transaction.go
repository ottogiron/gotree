package transaction

import (
	"github.com/ottogiron/gotree/api/backend/model"
)

type Type int

const (
	Remove = iota
	Add
	Update
)

type PersistHandler func(*T) error

type T struct {
	Type  Type
	Model *model.Tree
}

func NewTransaction(transactionType Type, tree *model.Tree) *T {
	return &T{transactionType, tree}
}
