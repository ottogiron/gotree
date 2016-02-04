package gotree

import (
	"net/url"

	"github.com/ottogiron/gotree/api"
)

type Repository struct {
}

func (c *Repository) Login() (api.Session, error) {
	return &Session{}, nil
}

func createRepository(url url.URL) (api.Repository, error) {
	return &Repository{}, nil
}
