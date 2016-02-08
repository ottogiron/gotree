package gotree

import (
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend"
	"github.com/ottogiron/gotree/api/backend/model"
)

type tree struct {
	*model.Tree
	kernel backend.Kernel
}

func NewTree(model *model.Tree, kernel backend.Kernel) api.Tree {
	return &tree{model, kernel}
}

func (t *tree) Name() string {
	return t.Tree.Name
}

func (t *tree) IsRoot() bool {
	return t.Path() == "/"
}

func (t *tree) Path() string {
	return t.Tree.Path
}

func (t *tree) Exists() bool {
	return t.Tree.Exists
}

func (t *tree) Parent() api.Tree {
	return nil
}

func (t *tree) Property() api.Property {
	return nil
}

func (t *tree) HasProperty(name string) bool {
	return false
}

func (t *tree) PropertyCount() int {
	return -1
}

func (t *tree) Properties() api.Properties {
	return nil
}

func (t *tree) Child(name string) error {
	return nil
}

func (t *tree) HasChild(name string) bool {
	return false
}

func (t *tree) ChildrenCount() int {
	return -1
}

func (t *tree) Children() []api.Tree {
	return nil
}

func (t *tree) Remove() error {
	return nil
}

func (t *tree) AddChild(name string) (api.Tree, error) {
	return t.kernel.AddChild(t.Path(), name)
}

func (t *tree) SetOrderableChildren(enable bool) {

}

func (t *tree) SetProperty(property api.Property) {

}

func (t *tree) SetPropertyValue(name string, value interface{}) error {
	return nil
}

func (t *tree) RemoveProperty(name string) {

}
