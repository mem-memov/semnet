package character

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character/node"
	"github.com/mem-memov/semnet/internal/class"
)

type entities struct {
	classes    *node.Classes
	bits       *node.Bits
	characters *node.Characters
	words      *node.Words
}

func newEntities(storage storage, classRepository *class.Repository, bitRepository *bit.Repository) *entities {
	return &entities{
		classes:    node.NewClasses(storage, classRepository),
		bits:       node.NewBits(storage, bitRepository),
		characters: node.NewCharacters(storage),
		words:      node.NewWords(storage),
	}
}

func (e *entities) create(classIdentifier uint, bitIdentifier uint, characterIdentifier uint, wordIdentifier uint) Entity {

	return newEntity(
		e.classes.Create(classIdentifier),
		e.bits.Create(bitIdentifier),
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
	)
}

func (e *entities) createAndAddClass(bitIdentifier uint, characterIdentifier uint, wordIdentifier uint) (Entity, error) {

	classNode, err := e.classes.CreateNew()
	if err != nil {
		return Entity{}, err
	}

	return newEntity(
		classNode,
		e.bits.Create(bitIdentifier),
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
	), nil
}

func (e *entities) createWithWord(wordIdentifier uint) (Entity, error) {

	wordNode := e.words.Create(wordIdentifier)

	characterIdentifier, err := wordNode.GetCharacter()
	if err != nil {
		return Entity{}, nil
	}

	characterNode := e.characters.Create(characterIdentifier)

	classIdentifier, bitIdentifier, wordIdentifierOfCharacter, err := characterNode.GetClassAndBitAndWord()
	if err != nil {
		return Entity{}, nil
	}

	if wordIdentifier != wordIdentifierOfCharacter {
		return Entity{}, fmt.Errorf("character has incorrect reference to word in character layer at character %d", characterIdentifier)
	}

	return e.create(classIdentifier, bitIdentifier, characterIdentifier, wordIdentifier), nil
}
