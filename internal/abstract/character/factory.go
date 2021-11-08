package character

import (
	"github.com/mem-memov/semnet/internal/abstract/bit"
	"github.com/mem-memov/semnet/internal/abstract/class"
)

type Factory interface {
	ProvideHeadEntity(classEntity class.Entity, bitEntity bit.Entity) (Entity, error)
	CreateTailEntity(classEntity class.Entity, bitEntity bit.Entity, previousCharacterEntity Entity) (Entity, error)
}
