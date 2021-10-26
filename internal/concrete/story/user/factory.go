package user

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractStoryFact "github.com/mem-memov/semnet/internal/abstract/story/fact"
	abstractStoryUser "github.com/mem-memov/semnet/internal/abstract/story/user"
)

type Factory struct {
	storage abstract.Storage
}

var _ abstractStoryUser.Factory = &Factory{}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateNewNode(factNode abstractStoryFact.Node) (abstractStoryUser.Node, error) {

	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	err = f.storage.SetReference(identifier, factNode.GetIdentifier())
	if err != nil {
		return nil, err
	}

	return newNode(identifier), nil
}
