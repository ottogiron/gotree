package api

import "io"

type Backend interface {
	io.Closer
	Tree(path string) (Tree, error)
	Move(sourcePath, destPath string) error
}
