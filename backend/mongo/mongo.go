package mongo

import "github.com/ottogiron/gotree/api"

type Mongo struct {
}

func New() (*Mongo, error) {
	return &Mongo{}, nil
}

func (m *Mongo) Close() error {
	return nil
}

func (m *Mongo) Move(sourcePath, destPath string) error {
	return nil
}

func (m *Mongo) Tree(path string) (api.Tree, error) {
	return nil, nil
}
