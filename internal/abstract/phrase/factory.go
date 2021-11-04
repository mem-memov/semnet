package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/word"
)

type Factory interface {
	ProvideEntity(classEntity class.Entity, wordEntity word.Entity) (Entity, error)
}
