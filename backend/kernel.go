package backend

import (
	"github.com/ottogiron/gotree"
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
)

type kernel struct {
	transactionManager transaction.Manager
	backend            backend.B
}

func NewKernel(backend backend.B) backend.Kernel {

	transactionManager := NewTransactionManager()
	return &kernel{
		transactionManager: transactionManager,
		backend:            backend,
	}
}

func (k *kernel) Open() error {
	return k.backend.Open()
}

func (k *kernel) Close() error {
	k.transactionManager.Persist(k.persistHandler())
	return k.backend.Close()
}

func (k *kernel) persistHandler() transaction.PersistHandler {

	persistHandler := func(transaction *transaction.T) error {
		return k.backend.Persist(transaction)
	}

	return persistHandler
}

func (m *kernel) Move(sourcePath, destPath string) error {
	return nil
}

func (m *kernel) Tree(path string) (api.Tree, error) {
	treeModel, err := m.backend.Tree(path)

	if err != nil {
		treeModel = &model.Tree{Exists: false}
	}

	tree := gotree.NewTree(treeModel, m)
	return tree, nil
}
