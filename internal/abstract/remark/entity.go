package remark

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/fact"
)

type Entity interface {
	GetClass() uint
	GetDetail() uint
	GetPosition() uint
	GetFact() uint
	PointToClass(class class.Entity) error
	PointToPosition(remark Entity) error
	PointToFact(fact fact.Aggregate) error
	FetchTargetFact(factRepository fact.Repository) (fact.Aggregate, error)
	CreateNextStoryFact(factRepository fact.Repository) (fact.Aggregate, error)
	HasNextRemark() (bool, error)
	GetNextRemark() (Entity, error)
	ToNextFact(fact uint) (Entity, error)
}
