package word

import (
	"github.com/mem-memov/semnet/internal/abstract/character"
	"github.com/mem-memov/semnet/internal/abstract/class"
)

type Storage interface {
	CreateEntity(classEntity class.Entity, characterEntity character.Entity) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByCharacter(character uint) (Entity, error)
	ReadEntityByWord(word uint) (Entity, error)
	ReadEntityByPhrase(phrase uint) (Entity, error)
}
