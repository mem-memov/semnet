package character

import (
	"github.com/mem-memov/semnet/internal/character/node"
)

type Entity struct {
	codeNode      node.Code
	characterNode node.Character
	wordNode      node.Word
}

func newEntity(codeNode node.Code, characterNode node.Character, wordNode node.Word) Entity {
	return Entity{
		codeNode:      codeNode,
		characterNode: characterNode,
		wordNode:      wordNode,
	}
}

func (e Entity) provideNext(codeValue rune, entities *entities) (Entity, error) {

	targetCharacters, err := e.characterNode.ReadTargets()
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetCharacter := range targetCharacters {
		bitIdentifier, characterIdentifier, err := targetCharacter.GetBitAndCharacter()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.create(bitIdentifier, targetCharacter.Identifier(), characterIdentifier)

		hasBitValue, err := entity.hasBitValue(codeValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasBitValue {
			return entity, nil
		}
	}

	// Provide new
	newBitNode, err := e.bitNode.NewBit(codeValue)
	if err != nil {
		return Entity{}, nil
	}

	newCodeNode, err := e.codeNode.NewCode(newBitNode)
	if err != nil {
		return Entity{}, nil
	}

	newCharacterNode, err := e.characterNode.NewCharacter(newCodeNode)
	if err != nil {
		return Entity{}, nil
	}

	return newEntity(newBitNode, newCodeNode, newCharacterNode), nil
}

