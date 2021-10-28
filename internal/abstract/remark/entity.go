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
	PointToFact(fact fact.Entity) error
	FetchTargetFact(factRepository fact.Repository) (fact.Entity, error)
	CreateNextStoryFact(factRepository fact.Repository) (fact.Entity, error)
	HasNextRemark() (bool, error)
	GetNextRemark() (Entity, error)
}