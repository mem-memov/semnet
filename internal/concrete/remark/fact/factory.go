package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemarkFact "github.com/mem-memov/semnet/internal/abstract/remark/fact"
)

type Factory struct {
	storage abstract.Storage
	factRepository abstractFact.Repository
}

var _ abstractRemarkFact.Factory = &Factory{}

func NewFactory(storage abstract.Storage, factRepository abstractFact.Repository) *Factory {
	return &Factory{
		storage: storage,
		factRepository: factRepository,
	}
}

func (f *Factory) CreateExistingNode(identifier uint) abstractRemarkFact.Node {
	return newNode(identifier, f.storage)
}

func (f *Factory) CreateNewNode() (abstractRemarkFact.Node, error) {
	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	_, err = f.factRepository.CreateNewFact(identifier)
	if err != nil {
		return nil, err
	}

	return newNode(identifier, f.storage), nil
}
