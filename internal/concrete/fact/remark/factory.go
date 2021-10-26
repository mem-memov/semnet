package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	abstractFactClass "github.com/mem-memov/semnet/internal/abstract/fact/class"
	abstractFactRemark "github.com/mem-memov/semnet/internal/abstract/fact/remark"
)

type Factory struct {
	storage abstract.Storage
}

var _ abstractFactRemark.Factory = &Factory{}

func NewFactory(storage abstract.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) CreateNewNode(
	remarkEntity abstractRemark.Entity,
	classNode abstractFactClass.Node,
) (abstractFactRemark.Node, error) {

	identifier, err := f.storage.Create()
	if err != nil {
		return nil, err
	}

	err = remarkEntity.Mark(identifier)
	if err != nil {
		return nil, err
	}

	err = f.storage.SetReference(classNode.GetIdentifier(), identifier)
	if err != nil {
		return nil, err
	}

	return newNode(identifier), nil
}
