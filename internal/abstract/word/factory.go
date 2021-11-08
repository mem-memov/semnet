package word

import (
	"github.com/mem-memov/semnet/internal/abstract/character"
	"github.com/mem-memov/semnet/internal/abstract/class"
)

type Factory interface {
	ProvideHeadEntity(classEntity class.Entity, characterEntity character.Aggregate) (Entity, error)
	CreateTailEntity(classEntity class.Entity, bitEntity character.Aggregate, previousWordEntity Entity) (Entity, error)
}
