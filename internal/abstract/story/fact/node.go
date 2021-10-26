package fact

import "github.com/mem-memov/semnet/internal/abstract/fact"

type Node interface {
	GetIdentifier() uint
	GetMarked(factEntity fact.Entity) error
}
