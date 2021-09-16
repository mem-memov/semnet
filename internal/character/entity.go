package character

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/character/node"
)

type Entity struct {
	classNode     node.Class
	bitNode       node.Bit
	characterNode node.Character
	wordNode      node.Word
}

func newEntity(classNode node.Class, bitNode node.Bit, characterNode node.Character, wordNode node.Word) Entity {
	return Entity{
		classNode:     classNode,
		bitNode:       bitNode,
		characterNode: characterNode,
		wordNode:      wordNode,
	}
}

func (e Entity) WordIdentifier() uint {
	return e.wordNode.Identifier()
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.wordNode.Mark(sourceIdentifier)
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	return e.wordNode.ProvideSingleTarget()
}

func (e Entity) provideNext(bitValue bool, entities *entities) (Entity, error) {

	targetCharacters, err := e.characterNode.ReadTargets()
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetCharacter := range targetCharacters {
		classIdentifier, bitIdentifier, wordIdentifier, err := targetCharacter.GetClassAndBitAndWord()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.create(classIdentifier, bitIdentifier, targetCharacter.Identifier(), wordIdentifier)

		hasBitValue, err := entity.hasBitValue(bitValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasBitValue {
			return entity, nil
		}
	}

	// Provide new
	newClassNode, err := e.classNode.NewClass()
	if err != nil {
		return Entity{}, nil
	}

	newBitNode, err := e.bitNode.NewBit(bitValue)
	if err != nil {
		return Entity{}, nil
	}

	newCharacterNode, err := e.characterNode.NewCharacter(newBitNode)
	if err != nil {
		return Entity{}, nil
	}

	newWordNode, err := e.wordNode.NewWord(newCharacterNode)
	if err != nil {
		return Entity{}, nil
	}

	return newEntity(newClassNode, newBitNode, newCharacterNode, newWordNode), nil
}

func (e Entity) hasBitValue(bit bool) (bool, error) {

	hasBitValue, err := e.bitNode.HasBitValue(bit)
	if err != nil {
		return false, err
	}

	return hasBitValue, nil
}

func (e Entity) findPrevious(entities *entities) (Entity, bool, error) {

	sourceCharacters, err := e.characterNode.ReadSources()
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourceCharacters) {
	case 0:
		return e, true, nil
	case 1:
		parentCharacter := sourceCharacters[0]

		classIdentifier, bitIdentifier, wordIdentifier, err := parentCharacter.GetClassAndBitAndWord()
		if err != nil {
			return Entity{}, false, nil
		}

		return entities.create(classIdentifier, bitIdentifier, parentCharacter.Identifier(), wordIdentifier), false, nil
	default:
		return Entity{}, false, fmt.Errorf("too many sources in character tree")
	}
}

func (e Entity) bitValue() (bool, error) {

	return e.bitNode.BitValue()
}

func (e Entity) String() string {
	return fmt.Sprintf("Character: %s %s %s\n", e.bitNode, e.characterNode, e.wordNode)
}
