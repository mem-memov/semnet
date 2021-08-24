package character

import (
	"fmt"
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

func (e *entities) createWithWord(wordIdentifier uint) (Entity, error) {

	wordNode := e.words.Create(wordIdentifier)

	characterIdentifier, err := wordNode.GetCharacter()
	if err != nil {
		return Entity{}, nil
	}

	characterNode := e.characters.Create(characterIdentifier)

	bitIdentifier, wordIdentifierOfCharacter, err := characterNode.GetBitAndWord()
	if err != nil {
		return Entity{}, nil
	}

	if wordIdentifier != wordIdentifierOfCharacter {
		return Entity{}, fmt.Errorf("character has incorrect reference to word in character layer at character %d", characterIdentifier)
	}

	return e.create(bitIdentifier, characterIdentifier, wordIdentifier), nil
}
