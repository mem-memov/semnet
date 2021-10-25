package class

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractRemarkClass "github.com/mem-memov/semnet/internal/abstract/remark/class"
)

// Factory structure creates class nodes for remark structures.
type Factory struct {
	repository abstractClass.Repository
	creator    *IdentifierCreator
}

var _ abstractRemarkClass.Factory = &Factory{}

// NewFactory creates a new instance of the factory.
// It takes a repository of class entities and a creator of new identifiers as parameters.
func NewFactory(repository abstractClass.Repository) *Factory {
	return &Factory{
		repository: repository,
		creator:    newIdentifierCreator(repository),
	}
}

// CreateExistingNode takes an identifier that is part of an existing remark.
// It wraps the identifier into a structure adding some useful methods to it.
func (f *Factory) CreateExistingNode(identifier uint) abstractRemarkClass.Node {
	return newNode(identifier, f.repository, f.creator)
}

// CreateNewNode helps to create a new remark providing a new node pointing to another node that is common for all remarks.
func (f *Factory) CreateNewNode() (abstractRemarkClass.Node, error) {
	newIdentifier, err := f.creator.CreateNewIdentifier()
	if err != nil {
		return nil, err
	}

	return newNode(newIdentifier, f.repository, f.creator), nil
}
