package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractStoryClass "github.com/mem-memov/semnet/internal/abstract/story/class"
	abstractStoryFact "github.com/mem-memov/semnet/internal/abstract/story/fact"
)

type Factory struct {
	storage abstract.Storage
}

var _ abstractStoryFact.Factory = &Factory{}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateNewNode(
	classNode abstractStoryClass.Node,
) (abstractStoryFact.Node, error) {

	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	err = f.storage.SetReference(classNode.GetIdentifier(), identifier)
	if err != nil {
		return nil, err
	}

	return newNode(identifier), nil
}
