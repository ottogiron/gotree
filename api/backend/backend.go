package backend

import (
	"io"

	"github.com/ottogiron/gotree/api"
)

type B interface {
	io.Closer
	Open() error
	Tree(path string) (api.Tree, error)
	Move(sourcePath, destPath string) error
}
