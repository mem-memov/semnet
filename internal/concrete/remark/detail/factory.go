package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractRemarkDetail "github.com/mem-memov/semnet/internal/abstract/remark/detail"
)

type Factory struct {
	storage    abstract.Storage
	repository abstractDetail.Repository
	creator    abstractRemarkDetail.IdentifierCreator
}

var _ abstractRemarkDetail.Factory = &Factory{}

func NewFactory(storage abstract.Storage, repository abstractDetail.Repository) *Factory {
	return &Factory{
		storage:    storage,
		repository: repository,
		creator: newIdentifierCreator(storage, repository),
	}
}

func (f *Factory) CreateExistingNode(identifier uint) abstractRemarkDetail.Node {
	return newNode(identifier, f.storage, f.repository, f.creator)
}

func (f *Factory) CreateNewNode(object string, property string) (abstractRemarkDetail.Node, error) {
	identifier, err := f.creator.CreateNewIdentifier(object, property)
	if err != nil {
		return Node{}, err
	}

	return newNode(identifier, f.storage, f.repository, f.creator), nil
}
