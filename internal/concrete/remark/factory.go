package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	abstractRemarkClass "github.com/mem-memov/semnet/internal/abstract/remark/class"
	abstractRemarkDetail "github.com/mem-memov/semnet/internal/abstract/remark/detail"
	abstractRemarkFact "github.com/mem-memov/semnet/internal/abstract/remark/fact"
	abstractRemarkPosition "github.com/mem-memov/semnet/internal/abstract/remark/position"
	"github.com/mem-memov/semnet/internal/concrete/remark/class"
	"github.com/mem-memov/semnet/internal/concrete/remark/detail"
	"github.com/mem-memov/semnet/internal/concrete/remark/fact"
	"github.com/mem-memov/semnet/internal/concrete/remark/position"
)

type Factory struct {
	classNodeFactory  abstractRemarkClass.Factory
	detailNodeFactory   abstractRemarkDetail.Factory
	positionNodeFactory abstractRemarkPosition.Factory
	factNodeFactory     abstractRemarkFact.Factory
}

var _ abstractRemark.Factory = &Factory{}

func newFactory(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	detailRepository abstractDetail.Repository,
) *Factory {
	return &Factory{
		classNodeFactory:    class.NewFactory(classRepository),
		detailNodeFactory:   detail.NewFactory(storage, detailRepository),
		positionNodeFactory: position.NewFactory(storage),
		factNodeFactory:     fact.NewFactory(storage),
	}
}

func (f *Factory) CreateExistingEntity(
	classIdentifier uint,
	detailIdentifier uint,
	remarkIdentifier uint,
	factIdentifier uint,
) abstractRemark.Entity {
	return newEntity(
		f.classNodeFactory.CreateExistingNode(classIdentifier),
		f.detailNodeFactory.CreateExistingNode(detailIdentifier),
		f.positionNodeFactory.CreateExistingNode(remarkIdentifier),
		f.factNodeFactory.CreateExistingNode(factIdentifier),
	)
}

func (f *Factory) CreateNewEntity(object string, property string) (abstractRemark.Entity, error) {
	classNode, err := f.classNodeFactory.CreateNewNode()
	if err != nil {
		return nil, err
	}

	detailNode, err := f.detailNodeFactory.CreateNewNode(object, property)
	if err != nil {
		return nil, err
	}

	positionNode, err := f.positionNodeFactory.CreateNewNode()
	if err != nil {
		return nil, err
	}

	factNode, err := f.factNodeFactory.CreateNewNode()
	if err != nil {
		return nil, err
	}

	entity := newEntity(classNode, detailNode, positionNode, factNode)

	return entity, nil
}
