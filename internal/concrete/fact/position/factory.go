package position

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractFactPosition "github.com/mem-memov/semnet/internal/abstract/fact/position"
	abstractFactRemark "github.com/mem-memov/semnet/internal/abstract/fact/remark"
)

type Factory struct {
	storage abstract.Storage
}

var _ abstractFactPosition.Factory = &Factory{}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateNewNode(remarkNode abstractFactRemark.Node) (abstractFactPosition.Node, error) {

	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	err = f.storage.SetReference(remarkNode.GetIdentifier(), identifier)
	if err != nil {
		return nil, err
	}

	return newNode(identifier), nil
}





