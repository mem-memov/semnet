package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/word"
)

type Storage interface {
	CreateEntity(classEntity class.Entity, wordEntity word.Aggregate) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByWord(word uint) (Entity, error)
	ReadEntityByPhrase(phrase uint) (Entity, error)
	ReadEntityByDetail(detail uint) (Entity, error)
}
