package gotree

import "github.com/ottogiron/gotree/api"

type Repository struct {
}

func (c *Repository) Login() (api.Session, error) {
	return &Session{}, nil
}

func CreateRepository(backend api.Backend) (api.Repository, error) {
	return &Repository{}, nil
}
