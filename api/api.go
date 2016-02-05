package api

import "io"

type PropertyType int

const (
	STRING = iota
	BINARY
	BOOL
	DATE
	FLOAT
)

type Repository interface {
	Login() (Session, error)
}

type Session interface {
	io.Closer
	Root() (Root, error)
}

type Root interface {
	Move(sourcePath, destPath string) error
	Tree(path string) (Tree, error)
	Session() (Session, error)
}

type Tree interface {
	Name() string
	IsRoot() bool
	Path() string
	Exists() bool
	Parent() Tree
	Property() Property
	HasProperty(name string) bool
	PropertyCount() int
	Properties() Properties
	Child(name string) error
	HasChild(name string) bool
	ChildrenCount() int
	Children() []Tree
	Remove() error
	AddChild(name string) Tree
	SetOrderableChildren(enable bool)
	SetProperty(property Property)
	SetPropertyValue(name string, value interface{}) error
	RemoveProperty(name string)
}

type Properties []Property
type Property interface {
	Type() PropertyType
	Tree() Tree
	String() string
	Binary() io.Reader
	Boolean() bool
	Date() string
	Float() float64
	Length() int
	Lenght() []int
	IsArray() bool
	SetValue(val interface{}) error
	Value() Value
	Values() Values
}

type Values []Value

type Value interface {
	Type() PropertyType
	String() string
	Binary() io.Reader
	Boolean() bool
	Date() string
	Float() float64
}
