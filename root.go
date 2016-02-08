package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type root struct {
	kernel  backend.Kernel
	session api.Session
}

func NewRoot(kernel backend.Kernel, session api.Session) api.Root {
	return &root{kernel, session}
}

func (r *root) Move(sourcePath, destPath string) error {
	return nil
}

func (r *root) Session() api.Session {
	return r.session
}

func (r *root) Tree(path string) (api.Tree, error) {
	return r.kernel.Tree(path)
}
