package backend

import (
	"github.com/ottogiron/gotree/api/backend"
	"github.com/ottogiron/gotree/api/backend/transaction"
	"github.com/ottogiron/gotree/backend/model"
)

type mapTransactionManager struct {
	operations map[transaction.Type]*model.Tree
}

func NewTransactionManager() transaction.Manager {
	return &mapTransactionManager{}
}

func (s *mapTransactionManager) Add(transactionType transaction.Type, tree *model.Tree) {
	s.operations[transactionType] = tree
}

func (s *mapTransactionManager) Persist(backend backend.B) error {
	return nil
}
