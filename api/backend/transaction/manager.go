package transaction

import "github.com/ottogiron/gotree/api/backend/model"

type Manager interface {
	Add(transactionType Type, tree *model.Tree)
	Persist(handler PersistHandler) error
	EnsureTreeWillExists(path string) bool
}
