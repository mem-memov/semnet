package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractRemarkClass "github.com/mem-memov/semnet/internal/abstract/remark/class"
	abstractRemarkDetail "github.com/mem-memov/semnet/internal/abstract/remark/detail"
	abstractRemarkFact "github.com/mem-memov/semnet/internal/abstract/remark/fact"
	abstractRemarkRemark "github.com/mem-memov/semnet/internal/abstract/remark/remark"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	"github.com/mem-memov/semnet/internal/concrete/remark/class"
	"github.com/mem-memov/semnet/internal/concrete/remark/detail"
	"github.com/mem-memov/semnet/internal/concrete/remark/fact"
	"github.com/mem-memov/semnet/internal/concrete/remark/remark"
)

type Factory struct {
	classNodeFactory  abstractRemarkClass.Factory
	detailNodeFactory abstractRemarkDetail.Factory
	remarkNodeFactory abstractRemarkRemark.Factory
	factNodeFactory   abstractRemarkFact.Factory
}

func NewFactory(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	detailRepository abstractDetail.Repository,
) *Factory {
	return &Factory{
		classNodeFactory:  class.NewFactory(classRepository),
		detailNodeFactory: detail.NewFactory(storage, detailRepository),
		remarkNodeFactory: remark.NewFactory(),
		factNodeFactory:   fact.NewFactory(),
	}
}

func (f *Factory) CreateExistingEntity(classIdentifier uint, detailIdentifier uint, remarkIdentifier uint, factIdentifier uint) abstractRemark.Entity {
	return newEntity(
		f.classNodeFactory.CreateExistingNode(classIdentifier),
		f.detailNodeFactory.CreateExistingNode(detailIdentifier),
		f.remarkNodeFactory.CreateExistingNode(remarkIdentifier),
		f.factNodeFactory.CreateExistingNode(factIdentifier),
	)
}

func (f *Factory) CreateNewEntity(detailIdentifier uint, remarkIdentifier uint, factIdentifier uint) abstractRemark.Entity {
	return newEntity(
		f.classNodeFactory.CreateNewNode(),
		f.detailNodeFactory.CreateExistingNode(detailIdentifier),
		f.remarkNodeFactory.CreateExistingNode(remarkIdentifier),
		f.factNodeFactory.CreateExistingNode(factIdentifier),
	)
}
