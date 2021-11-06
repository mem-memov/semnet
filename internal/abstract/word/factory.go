package word

import (
	"github.com/mem-memov/semnet/internal/abstract/character"
	"github.com/mem-memov/semnet/internal/abstract/class"
)

type Factory interface {
	ProvideFirstEntity(classEntity class.Entity, characterEntity character.Entity) (Entity, error)
}
