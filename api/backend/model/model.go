package model

type Tree struct {
	Name       string
	Path       string
	Properties map[string]interface{}
	Exists     bool
}

func NewTree(path string, exists bool) *Tree {
	return &Tree{Path: path, Exists: exists}
}
