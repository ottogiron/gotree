package mongo

import (
	"fmt"
	"strings"

	"github.com/ottogiron/gotree"
	"github.com/ottogiron/gotree/api"
	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
	"github.com/ottogiron/gotree/backend"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongo struct {
	transactionManager   transaction.Manager
	session              *mgo.Session
	connectionString     string
	repositoryCollection string
	repositoryDB         string
	sessionTrees         []*model.Tree
}

func New(connStrings, repositoryDB, repositoryCollection string) *mongo {
	transactionManager := backend.NewTransactionManager()
	return &mongo{
		transactionManager:   transactionManager,
		connectionString:     connStrings,
		repositoryDB:         repositoryDB,
		repositoryCollection: repositoryCollection}
}

func (m *mongo) Open() error {
	var err error
	m.session, err = mgo.Dial(m.connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongo) Close() error {

	defer m.session.Close()
	err := m.transactionManager.Persist(persistHandler)
	if err != nil {
		return err
	}
	return nil
}

func persistHandler(transaction.T) error {
	return nil
}

func (m *mongo) Move(sourcePath, destPath string) error {
	return nil
}

func (m *mongo) Tree(path string) (api.Tree, error) {
	c := m.collection()

	path, err := transformTomongoPath(path)

	if err != nil {
		return nil, err
	}

	result := &model.Tree{}
	err = c.Find(bson.M{"path": path}).One(result)

	if err != nil {
		result.Exists = false
	}

	tree := gotree.NewTree(result, m)
	return tree, nil
}

func (m *mongo) collection() *mgo.Collection {
	collection := m.session.DB(m.repositoryDB).C(m.repositoryCollection)
	return collection
}

func transformTomongoPath(path string) (string, error) {
	if !isValidPath(path) {
		return "", fmt.Errorf("Invalid path %s", path)
	}
	transformedPath := strings.Replace(path, "/", ",", -1)

	return transformedPath, nil
}

func isValidPath(path string) bool {
	return true
}
