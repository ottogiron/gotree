package mongo

import (
	"testing"

	"github.com/ottogiron/gotree"
)

func TestGetTree(t *testing.T) {

	mongoBackend, err := New()

	if err != nil || mongoBackend == nil {
		t.Fatal("Error when connecting to the mongo database")
	}

	repository, err := gotree.CreateRepository(mongoBackend)

	if err != nil || repository == nil {
		t.Fatal("Error when creating repository")
	}

	session, err := repository.Login()

	if err != nil && session == nil {
		t.Fatal("Could not create a session")
	}

	root, err := session.Root()

	if err != nil || root == nil {
		t.Fatal("Could not get the root of the repository")
	}

	rootTree, err := root.Tree("/")

	if err != nil || rootTree == nil {
		t.Fatal("Couldn't get the root tree")
	}

}
