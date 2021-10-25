package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	"github.com/mem-memov/semnet/internal/abstract/detail"
)

// IdentifierCreator structure is a service for several other structures that need it.
// This service creates new detail identifiers for remark entities.
type IdentifierCreator struct {
	storage    abstract.Storage
	repository detail.Repository
}

func newIdentifierCreator(storage abstract.Storage, repository detail.Repository) *IdentifierCreator {
	return &IdentifierCreator{
		storage: storage,
		repository: repository,
	}
}

// CreateNewIdentifier creates a new detail identifier to be used as part of a remark cluster.
func (i *IdentifierCreator) CreateNewIdentifier(object string, property string) (uint, error) {

	entity, err := i.repository.Provide(object, property)
	if err != nil {
		return 0, err
	}

	newIdentifier, err := i.storage.Create()
	if err != nil {
		return 0, err
	}

	err = entity.AddRemark(newIdentifier)
	if err != nil {
		return 0, err
	}

	return newIdentifier, nil
}