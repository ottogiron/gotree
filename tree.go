package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/backend/model"
)

type Tree struct {
	*model.Tree
}

func (t *Tree) Name() string {
	return t.Tree.Name
}

func (t *Tree) IsRoot() bool {
	return false
}

func (t *Tree) Path() string {
	return t.Tree.Path
}

func (t *Tree) Exists() bool {
	return true
}

func (t *Tree) Parent() api.Tree {
	return nil
}

func (t *Tree) Property() api.Property {
	return nil
}

func (t *Tree) HasProperty(name string) bool {
	return false
}

func (t *Tree) PropertyCount() int {
	return -1
}

func (t *Tree) Properties() api.Properties {
	return nil
}

func (t *Tree) Child(name string) error {
	return nil
}

func (t *Tree) HasChild(name string) bool {
	return false
}

func (t *Tree) ChildrenCount() int {
	return -1
}

func (t *Tree) Children() []api.Tree {
	return nil
}

func (t *Tree) Remove() error {
	return nil
}

func (t *Tree) AddChild(name string) api.Tree {
	return nil
}

func (t *Tree) SetOrderableChildren(enable bool) {

}

func (t *Tree) SetProperty(property api.Property) {

}

func (t *Tree) SetPropertyValue(name string, value interface{}) error {
	return nil
}

func (t *Tree) RemoveProperty(name string) {

}
