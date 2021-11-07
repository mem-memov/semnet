package bit

import (
	"github.com/mem-memov/semnet/internal/abstract"
	"github.com/mem-memov/semnet/internal/concrete/bit/node"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type entities struct {
	classes    *node.Classes
	characters *node.Characters
}

func newEntities(storage abstract.Storage, classRepository *class.Repository) *entities {
	return &entities{
		classes:    node.NewClasses(storage, classRepository),
		characters: node.NewCharacters(storage),
	}
}

func (e *entities) create(value bool, classIdentifier uint, characterIdentifier uint) Entity {

	return newEntity(
		value,
		e.classes.Create(classIdentifier),
		e.characters.Create(characterIdentifier),
	)
}
