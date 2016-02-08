package backend

import (
	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
)

type SliceTransactionManager struct {
	transactions []*transaction.T
}

func NewTransactionManager() transaction.Manager {
	return &SliceTransactionManager{}
}

func (s *SliceTransactionManager) Add(transactionType transaction.Type, tree *model.Tree) {
	t := transaction.NewTransaction(transactionType, tree)
	s.transactions = append(s.transactions, t)
}

func (s *SliceTransactionManager) EnsureTreeWillExists(path string) bool {
	treeTransactionTypes := []transaction.Type{}

	for _, t := range s.transactions {
		if t.Model.Path == path {
			treeTransactionTypes = append(treeTransactionTypes, t.Type)
		}
	}
	//Check the transaction list backwards if the last added transaction is Add
	var ty transaction.Type
	for i := len(treeTransactionTypes) - 1; i >= 0; i-- {
		ty = treeTransactionTypes[i]
		if ty == transaction.Remove {
			return false
		} else if ty == transaction.Add {
			return true
		}
	}
	return false
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
