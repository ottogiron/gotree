package gotree

import "github.com/ottogiron/gotree/api"

type Session struct {
}

func (s *Session) Close() error {
	return nil
}

func (s *Session) Root() (api.Root, error) {
	return &Root{}, nil
}
