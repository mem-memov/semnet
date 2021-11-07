package character

import (
	"github.com/mem-memov/semnet/internal/abstract/bit"
	"github.com/mem-memov/semnet/internal/abstract/class"
)

type Storage interface {
	CreateEntity(classEntity class.Entity, bitEntity bit.Entity) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByBit(bit uint) (Entity, error)
	ReadEntityByCharacter(character uint) (Entity, error)
	ReadEntityByWord(word uint) (Entity, error)
}
