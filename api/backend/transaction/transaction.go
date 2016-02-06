package transaction

import (
	"github.com/ottogiron/gotree/api/backend"
	"github.com/ottogiron/gotree/backend/model"
)

type Type int

const (
	Remove = iota
	Add
	Update
)

type Manager interface {
	Add(transactionType Type, tree *model.Tree)
	Persist(backend backend.B) error
}
