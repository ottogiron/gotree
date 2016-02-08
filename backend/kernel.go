package backend

import (
	"errors"
	"fmt"
	"path"

	"github.com/ottogiron/gotree"
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
)

type kernel struct {
	transactionManager transaction.Manager
	backend            backend.B
	closed             bool
}

func NewKernel(backend backend.B) backend.Kernel {

	transactionManager := NewTransactionManager()
	return &kernel{
		transactionManager: transactionManager,
		backend:            backend,
		closed:             false,
	}
}

func (k *kernel) Open() error {
	return k.backend.Open()
}

func (k *kernel) Close() error {

	if !k.closed {
		k.transactionManager.Persist(k.persistHandler())
		if err := k.backend.Close(); err != nil {
			return fmt.Errorf("Error when trying to close session %s", err)
		}
		k.closed = true
		return nil
	}

	return errors.New("Error when trying to close session: already closed")
}

func (k *kernel) Tree(path string) (api.Tree, error) {
	treeModel, err := k.backend.Tree(path)

	if err != nil {
		return nil, err
	}
	tree := gotree.NewTree(treeModel, k)
	return tree, nil
}

func (k *kernel) AddChild(parentPath, name string) (api.Tree, error) {

	parentWillExists := k.transactionManager.EnsureTreeWillExists(parentPath)
	if !parentWillExists {
		parentModel, err := k.backend.Tree(parentPath)
		if !parentModel.Exists || err != nil {
			return nil, fmt.Errorf("Parent %s does not exis, please create the parent tree first", parentPath)
		}
	}
	childPath := path.Join(parentPath, name)
	childModel := model.NewTree(childPath, false)
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
