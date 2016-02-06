package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type Session struct {
	backend backend.B
}

func NewSession(backend backend.B) api.Session {
	return &Session{backend}
}

func (s *Session) Close() error {
	return s.backend.Close()
}

func (s *Session) Root() api.Root {
	return NewRoot(s.backend, s)
}
