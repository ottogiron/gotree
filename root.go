package gotree

import "github.com/ottogiron/gotree/api"

type Root struct {
}

func (r *Root) Move(sourcePath, destPath string) error {
	return nil
}

func (r *Root) Session() (api.Session, error) {
	return nil, nil
}

func (r *Root) Tree(path string) (api.Tree, error) {
	return nil, nil
}
