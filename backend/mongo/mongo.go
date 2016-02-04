package mongo

import (
	"github.com/ottogiron/gotree/api"
	"gopkg.in/mgo.v2"
)

type Mongo struct {
	session          *mgo.Session
	connectionString string
}

func New(connStrings string) *Mongo {

	return &Mongo{connectionString: connStrings}
}

func (m *Mongo) Open() error {
	var err error
	m.session, err = mgo.Dial(m.connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) Close() error {
	m.session.Close()
	return nil
}

func (m *Mongo) Move(sourcePath, destPath string) error {
	return nil
}

func (m *Mongo) Tree(path string) (api.Tree, error) {
	return nil, nil
}
