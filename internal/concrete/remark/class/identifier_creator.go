package class

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
)

// IdentifierCreator structure is a service for several other structures that need it.
// This service creates new class identifiers for position entities.
// Each position cluster must have one node pointing to the node that is common to all remarks.
// This creates a cluster of remarks inside the all encompassing semantic graph.
type IdentifierCreator struct {
	repository abstractClass.Repository
}

func newIdentifierCreator(repository abstractClass.Repository) *IdentifierCreator {
	return &IdentifierCreator{
		repository: repository,
	}
}

// CreateNewIdentifier created new class identifier to be used as part of a position cluster.
func (i *IdentifierCreator) CreateNewIdentifier() (uint, error) {
	classEntity, err := i.repository.ProvideEntity()
	if err != nil {
		return 0, err
	}

	newIdentifier, err := classEntity.CreateRemark()
	if err != nil {
		return 0, err
	}

	return newIdentifier, nil
}