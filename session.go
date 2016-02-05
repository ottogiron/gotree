package gotree

import "github.com/ottogiron/gotree/api"

type Session struct {
	backend api.Backend
}

func (s *Session) Close() error {
	return s.backend.Close()
}

func (s *Session) Root() (api.Root, error) {
	return &Root{s.backend}, nil
}
