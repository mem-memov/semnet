package detail

import (
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Storage interface {
	CreateEntity(classEntity class.Entity, objectPhrase phrase.Entity, propertyPhrase phrase.Entity) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByPhrase(phrase uint) (Entity, error)
	ReadEntityByRemark(remark uint) (Entity, error)
}
