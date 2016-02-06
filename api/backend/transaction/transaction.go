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

type Manager interface {
	Add(transactionType Type, tree *model.Tree)
	Persist(handler PersistHandler) error
}

type T struct {
	transactionType Type
	tree            *model.Tree
}

func NewTransaction(transactionType Type, tree *model.Tree) *T {
	return &T{transactionType, tree}
}
