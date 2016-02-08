package mongo

import (
	"os"
	"testing"

	"github.com/ottogiron/gotree"
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/backend"
)

func createSession(t *testing.T) (api.Session, error) {
	hosts := os.Getenv("MONGO_TEST_HOSTS")
	repositoryCollection := "monto_tree_test"
	repositoryDB := "db_test"
	mongoBackend := New(hosts, repositoryDB, repositoryCollection)
	kernel := backend.NewKernel(mongoBackend)
	repository := gotree.NewRepository(kernel)

	session, err := repository.Login()
	return session, err
}

func TestGetRootTree(t *testing.T) {
	t.Parallel()

	session, err := createSession(t)

	if err != nil {
		t.Fatal("Could not create a session")
	}

	defer session.Close()

	root := session.Root()

	rootTree, err := root.Tree("/")

	if err != nil {
		t.Fatalf("Couldn't get the root tree: %s", err)
	}

	exists := rootTree.Exists()

	if exists {
		t.Fatal("Tree should not exist")
	}

}

func TestAddChild(t *testing.T) {
	t.Parallel()

	session, _ := createSession(t)

	defer session.Close()

	root := session.Root()
	rootTree, err := root.Tree("/")

	if err != nil {
		t.Fatalf("Couldn't get the root tree: %s", err)
	}

	testChild := rootTree.AddChild("testChild")

	if testChild == nil {
		t.Fatal("Test chils shoud not be nil")
	}

	if testChild.Path() != "/testChild" {
		t.Fatal("Test path should be /testChild")
	}

}
