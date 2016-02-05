package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/backend/model"
)

type Root struct {
	backend api.Backend
}

func (r *Root) Move(sourcePath, destPath string) error {
	return nil
}

func (r *Root) Session() (api.Session, error) {
	return nil, nil
}

func (r *Root) Tree(path string) (api.Tree, error) {
	if path == "/" {
		treeModel := &model.Tree{Name: "root", Path: "/"}
		return &Tree{treeModel, r.backend}, nil
	}
	return r.backend.Tree(path)
}
