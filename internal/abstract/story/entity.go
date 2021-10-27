package story

import "github.com/mem-memov/semnet/internal/abstract/fact"

type Entity interface {
	GetClass() uint
	GetFact() uint
	GetUser() uint
	PointToFact(fact fact.Entity) error
}
