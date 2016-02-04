package api

import "io"

type Backend interface {
	io.Closer
	Open() error
	Tree(path string) (Tree, error)
	Move(sourcePath, destPath string) error
}
