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

	if !exists {
		t.Fatal("Tree should exists")
	}

}

func TestAddChild(t *testing.T) {
	t.Parallel()

	session, _ := createSession(t)

	root := session.Root()
	rootTree, err := root.Tree("/")

	if err != nil {
		t.Fatalf("Couldn't get the root tree: %s", err)
	}

	if rootTree.Path() != "/" {
		t.Fatalf("Root tree path should be %s", rootTree.Path())
	}

	if !rootTree.Exists() {
		t.Fatal("Root tree / should exists")
	}
	childName := "testChild"
	testChild, err := rootTree.AddChild("testChild")

	if err != nil {
		t.Fatalf("Test chils shoud not be nil: %s", err)
	}

	if testChild.Exists() {
		t.Fatalf("Test child %s should not exist", testChild.Path())
	}

	if testChild.Path() != "/testChild" {
		t.Fatalf("Test path should be /testChild and it is: %s", testChild.Path())
	}

	session.Close()

	//Creat a new session to ensure the tree is created and exists
	secondSession, _ := createSession(t)
	childPath := "/" + childName
	root = secondSession.Root()
	savedChild, err := root.Tree(childPath)

	if err != nil {
		t.Fatalf("Child search %s should not return an error:%s", savedChild.Path(), err)
	}

	if !savedChild.Exists() {
		t.Fatalf("Child %s should exist", savedChild.Path())
	}
	secondSession.Close()
}
