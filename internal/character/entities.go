package character

import (
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character/node"
)

type entities struct {
	bits       *node.Bits
	characters *node.Characters
	words      *node.Words
}

func newEntities(storage storage, bitRepository *bit.Repository) *entities {
	return &entities{
		bits:       node.NewBits(storage, bitRepository),
		characters: node.NewCharacters(storage),
		words:      node.NewWords(storage),
	}
}

func (e *entities) create(bitIdentifier uint, characterIdentifier uint, wordIdentifier uint) Entity {

	return newEntity(
		e.bits.Create(bitIdentifier),
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
	)
}
