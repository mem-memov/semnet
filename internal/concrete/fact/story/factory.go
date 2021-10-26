package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractFactPosition "github.com/mem-memov/semnet/internal/abstract/fact/position"
	abstractFactStory "github.com/mem-memov/semnet/internal/abstract/fact/story"
)

type Factory struct {
	storage abstract.Storage
}

var _ abstractFactStory.Factory = &Factory{}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateNewNode(positionNode abstractFactPosition.Node) (abstractFactStory.Node, error) {

	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	err = f.storage.Connect(positionNode.GetIdentifier(), identifier)
	if err != nil {
		return nil, err
	}

	return newNode(identifier, f.storage), nil
}