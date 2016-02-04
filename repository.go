package gotree

import "github.com/ottogiron/gotree/api"

type Repository struct {
	backend api.Backend
}

func (c *Repository) Login() (api.Session, error) {
	if err := c.backend.Open(); err != nil {
		return nil, err
	}
	return &Session{c.backend}, nil
}

func CreateRepository(backend api.Backend) (api.Repository, error) {
	return &Repository{backend}, nil
}
