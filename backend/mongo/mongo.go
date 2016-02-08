package mongo

import (
	"fmt"
	"strings"

	"github.com/ottogiron/gotree/api/backend/model"
	"github.com/ottogiron/gotree/api/backend/transaction"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongo struct {
	session              *mgo.Session
	connectionString     string
	repositoryCollection string
	repositoryDB         string
	sessionTrees         []*model.Tree
}

func New(connStrings, repositoryDB, repositoryCollection string) *mongo {

	return &mongo{
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
	m.session.Close()
	return nil
}

func (m *mongo) Move(sourcePath, destPath string) error {
	return nil
}

func (m *mongo) Tree(path string) (*model.Tree, error) {
	c := m.collection()
	path, err := mongoPath(path)
	if err != nil {
		return nil, err
	}

	result := model.NewTree(path, true)

	err = c.Find(bson.M{"path": path}).One(result)

	if err != nil {
		result.Exists = false
	}

	return result, nil
}

func (m *mongo) Persist(t *transaction.T) error {
	model := t.Model
	switch t.Type {
	case transaction.Add:
		if !model.Exists {
			m.save(model)
			return nil
		}
	case transaction.Update:
		if !model.Exists {
			m.save(model)
			return nil
		}
		mPath, err := mongoPath(model.Path)
		if err != nil {
			return err
		}
		m.collection().Update(bson.M{"path": mPath}, model)

	case transaction.Remove:
		//Expensive operation to remove many children
	}
	return nil
}

func (m *mongo) save(treeModel *model.Tree) error {
	err := m.collection().Insert(treeModel)
	return err
}

func (m *mongo) collection() *mgo.Collection {
	collection := m.session.DB(m.repositoryDB).C(m.repositoryCollection)
	return collection
}

func mongoPath(path string) (string, error) {
	if !isValidPath(path) {
		return "", fmt.Errorf("Invalid path %s", path)
	}
	transformedPath := strings.Replace(path, "/", ",", -1)

	return transformedPath, nil
}

func isValidPath(path string) bool {
	return true
}
