package code

import (
	"github.com/mem-memov/semnet/internal/code/node"
)

type Entity struct {
	bitNode       node.Bit
	codeNode      node.Code
	characterNode node.Character
}

func newEntity(bitNode node.Bit, codeNode node.Code, characterNode node.Character) Entity {
	return Entity{
		bitNode:       bitNode,
		codeNode:      codeNode,
		characterNode: characterNode,
	}
}

func (e Entity) provideNext(bit bool, entities *entities) (Entity, error) {

	targetCodes, err := e.codeNode.ReadTargets()
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetCode := range targetCodes {
		bitIdentifier, characterIdentifier, err := targetCode.GetBitAndCharacter()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.create(bitIdentifier, targetCode.Identifier(), characterIdentifier)

		hasBitValue, err := entity.hasBitValue(bit)
		if err != nil {
			return Entity{}, nil
		}

		if hasBitValue {
			return entity, nil
		}
	}

	// Provide new
	newBitNode, err := e.bitNode.NewBit()
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

func (e Entity) hasBitValue(bit bool) (bool, error) {

	hasBitValue, err := e.bitNode.HasBitValue(bit)
	if err != nil {
		return false, err
	}

	return hasBitValue, nil
}
