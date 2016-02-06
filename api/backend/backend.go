package backend

import (
	"io"

	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
)

type B interface {
	io.Closer
	Open() error
	Tree(path string) (*model.Tree, error)
	Move(sourcePath, destPath string) error
	Persist(transaction *transaction.T) error
}

type Kernel interface {
	io.Closer
	Open() error
	Tree(path string) (api.Tree, error)
	Move(sourcePath, destPath string) error
}
