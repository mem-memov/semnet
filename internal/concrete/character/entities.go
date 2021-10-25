package character

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/concrete/bit"
	node2 "github.com/mem-memov/semnet/internal/concrete/character/node"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type entities struct {
	classes    *node2.Classes
	bits       *node2.Bits
	characters *node2.Characters
	words      *node2.Words
}

func newEntities(storage storage, classRepository *class.Repository, bitRepository *bit.Repository) *entities {
	return &entities{
		classes:    node2.NewClasses(storage, classRepository),
		bits:       node2.NewBits(storage, bitRepository),
		characters: node2.NewCharacters(storage),
		words:      node2.NewWords(storage),
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
