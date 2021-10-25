package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
)

type Factory struct {
	storage abstract.Storage
}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateExistingNode(identifier uint) Node {
	return newNode(identifier, f.storage)
}
