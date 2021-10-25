package bit

import (
	node2 "github.com/mem-memov/semnet/internal/concrete/bit/node"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type entities struct {
	classes    *node2.Classes
	characters *node2.Characters
}

func newEntities(storage storage, classRepository *class.Repository) *entities {
	return &entities{
		classes:    node2.NewClasses(storage, classRepository),
		characters: node2.NewCharacters(storage),
	}
}

func (e *entities) create(value bool, classIdentifier uint, characterIdentifier uint) Entity {

	return newEntity(
		value,
		e.classes.Create(classIdentifier),
		e.characters.Create(characterIdentifier),
	)
}
