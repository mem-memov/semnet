package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/word"
)

type Factory interface {
	ProvideFirstEntity(classEntity class.Entity, wordEntity word.Aggregate) (Entity, error)
}
