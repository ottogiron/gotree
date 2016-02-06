package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type repository struct {
	backend backend.B
}

func NewRepository(backend backend.B) api.Repository {
	return &repository{backend}
}

func (c *repository) Login() (api.Session, error) {
	if err := c.backend.Open(); err != nil {
		return nil, err
	}
	return NewSession(c.backend), nil
}
