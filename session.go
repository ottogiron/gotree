package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type Session struct {
	kernel backend.Kernel
}

func NewSession(kernel backend.Kernel) api.Session {
	return &Session{kernel}
}

func (s *Session) Close() error {
	return s.kernel.Close()
}

func (s *Session) Root() api.Root {
	return NewRoot(s.kernel, s)
}
