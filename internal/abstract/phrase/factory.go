package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/word"
)

type Factory interface {
	ProvideHeadEntity(classEntity class.Entity, wordEntity word.Aggregate) (Entity, error)
	CreateTailEntity(classEntity class.Entity, wordAggregate word.Aggregate, previousPhraseEntity Entity) (Entity, error)
}
