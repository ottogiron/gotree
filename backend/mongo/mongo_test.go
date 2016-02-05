package mongo

import (
	"os"
	"testing"

	"github.com/ottogiron/gotree"
)

func TestGetTree(t *testing.T) {

	hosts := os.Getenv("MONGO_TEST_HOSTS")
	mongoBackend := New(hosts)
	repository, err := gotree.CreateRepository(mongoBackend)

	if err != nil || repository == nil {
		t.Fatal("Error when creating repository")
	}

	session, err := repository.Login()

	if err != nil {
		t.Fatal("Could not create a session")
	}

	defer session.Close()

	root, err := session.Root()

	if err != nil || root == nil {
		t.Fatal("Could not get the root of the repository")
	}

	rootTree, err := root.Tree("/")

	if err != nil || rootTree == nil {
		t.Fatal("Couldn't get the root tree")
	}

}
