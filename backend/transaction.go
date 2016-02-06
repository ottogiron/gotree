package backend

import (
	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
)

type SliceTransactionManager struct {
	transactions []transaction.T
}

func NewTransactionManager() transaction.Manager {
	return &SliceTransactionManager{}
}

func (s *SliceTransactionManager) Add(transactionType transaction.Type, tree *model.Tree) {
	t := transaction.NewTransaction(transactionType, tree)
	s.transactions = append(s.transactions, t)
}

func (s *SliceTransactionManager) Persist(handler transaction.PersistHandler) error {

	for _, t := range s.transactions {
		err := handler(t)
		if err != nil {
			return err
		}
	}

	return nil
}
