package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type Root struct {
	backend backend.B
	session api.Session
}

func NewRoot(backend backend.B, session api.Session) api.Root {
	return &Root{backend, session}
}

func (r *Root) Move(sourcePath, destPath string) error {
	return nil
}

func (r *Root) Session() api.Session {
	return r.session
}

func (r *Root) Tree(path string) (api.Tree, error) {
	return r.backend.Tree(path)
}
