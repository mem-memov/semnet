package remark

import (
	"github.com/mem-memov/semnet/internal/abstract/fact"
)

type Entity interface {
	GetClass() uint
	GetDetail() uint
	GetPosition() uint
	GetFact() uint
	PointToPosition(remark Entity) error
	PointToFact(fact fact.Aggregate) error
	GetTargetFact() (uint, error)
	GetSourceDetail() (uint, error)
	CreateNextStoryFact(factRepository fact.Repository) (fact.Aggregate, error)
	HasNextRemark() (bool, error)
	GetNextRemark() (uint, error)
	GetNextFact(fact uint) (Entity, error)
}
