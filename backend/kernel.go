package backend

import (
	"fmt"

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

func (k *kernel) AddChild(parentPath, childPath string) (api.Tree, error) {

	addTransactionExists := k.transactionManager.AddTransactionExist(parentPath)
	if !addTransactionExists {
		_, err := k.backend.Tree(parentPath)
		if err != nil {
			return nil, fmt.Errorf("Parent %s does not exis, please create the parent tree first", parentPath)
		}

	}
	childModel := model.NewTree(childPath)
	k.transactionManager.Add(transaction.Add, childModel)
	tree := gotree.NewTree(childModel, k)
	return tree, nil
}

func (k *kernel) persistHandler() transaction.PersistHandler {

	persistHandler := func(transaction *transaction.T) error {
		return k.backend.Persist(transaction)
	}

	return persistHandler
}

func (k *kernel) Move(sourcePath, destPath string) error {
	return nil
}

func (k *kernel) Tree(path string) (api.Tree, error) {
	treeModel, err := k.backend.Tree(path)

	if err != nil {
		treeModel = &model.Tree{Exists: false}
	}

	tree := gotree.NewTree(treeModel, k)
	return tree, nil
}
