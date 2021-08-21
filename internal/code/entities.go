package code

import (
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/code/node"
)

type entities struct {
	bits       *node.Bits
	codes      *node.Codes
	characters *node.Characters
}

func newEntities(storage storage, bitRepository *bit.Repository) *entities {
	return &entities{
		bits:       node.NewBits(storage, bitRepository),
		codes:      node.NewCodes(storage),
		characters: node.NewCharacters(storage),
	}
}

func (e *entities) create(bitIdentifier uint, codeIdentifier uint, characterIdentifier uint) Entity {
	return newEntity(
		e.bits.Create(bitIdentifier),
		e.codes.Create(codeIdentifier),
		e.characters.Create(characterIdentifier),
	)
}
