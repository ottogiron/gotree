package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type session struct {
	kernel backend.Kernel
}

func NewSession(kernel backend.Kernel) api.Session {
	return &session{kernel}
}

func (s *session) Close() error {
	return s.kernel.Close()
}

func (s *session) Root() api.Root {
	return NewRoot(s.kernel, s)
}
