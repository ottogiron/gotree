package mongo

import (
	"os"
	"testing"

	"github.com/ottogiron/gotree"
	"github.com/ottogiron/gotree/api"
)

func createSession(t *testing.T) (api.Session, error) {
	hosts := os.Getenv("MONGO_TEST_HOSTS")
	repositoryCollection := "monto_tree_test"
	repositoryDB := "db_test"
	mongoBackend := New(hosts, repositoryDB, repositoryCollection)

	repository := gotree.NewRepository(mongoBackend)

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
