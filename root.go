package gotree

import "github.com/ottogiron/gotree/api"

type Root struct {
	session api.Session
	tree    api.Tree
}

func (r *Root) Move(sourcePath, destPath string) error {
	return nil
}

func (r *Root) Session() (api.Session, error) {
	return r.session, nil
}

func (r *Root) Tree(path string) (api.Tree, error) {
	return r.tree, nil
}
