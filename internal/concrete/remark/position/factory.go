package position

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractRemarkPosition "github.com/mem-memov/semnet/internal/abstract/remark/position"
)

type Factory struct {
	storage abstract.Storage
}

var _ abstractRemarkPosition.Factory = &Factory{}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateExistingNode(identifier uint) abstractRemarkPosition.Node {
	return newNode(identifier, f.storage)
}

func (f *Factory) CreateNewNode() (abstractRemarkPosition.Node, error) {
	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	return newNode(identifier, f.storage), nil
}
