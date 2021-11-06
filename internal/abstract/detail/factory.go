package detail

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Factory interface {
	ProvideEntity(classEntity class.Entity, objectPhrase phrase.Aggregate, propertyPhrase phrase.Aggregate) (Entity, error)
}
