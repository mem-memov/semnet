package story

import "github.com/mem-memov/semnet/internal/abstract/fact"

type Repository interface {
	CreateNewEntity(factEntity fact.Entity) (Entity, error)
}
