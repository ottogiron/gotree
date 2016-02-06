package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type Root struct {
	kernel  backend.Kernel
	session api.Session
}

func NewRoot(kernel backend.Kernel, session api.Session) api.Root {
	return &Root{kernel, session}
}

func (r *Root) Move(sourcePath, destPath string) error {
	return nil
}

func (r *Root) Session() api.Session {
	return r.session
}

func (r *Root) Tree(path string) (api.Tree, error) {
	return r.kernel.Tree(path)
}
