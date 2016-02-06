package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
)

type repository struct {
	kernel backend.Kernel
}

func NewRepository(kernel backend.Kernel) api.Repository {
	return &repository{kernel}
}

func (c *repository) Login() (api.Session, error) {
	if err := c.kernel.Open(); err != nil {
		return nil, err
	}
	return NewSession(c.kernel), nil
}
