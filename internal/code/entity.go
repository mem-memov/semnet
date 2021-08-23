package code

import (
	"fmt"
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

func (e Entity) provideNext(bitValue bool, entities *entities) (Entity, error) {

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

		hasBitValue, err := entity.hasBitValue(bitValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasBitValue {
			return entity, nil
		}
	}

	// Provide new
	newBitNode, err := e.bitNode.NewBit(bitValue)
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

func (e Entity) findPrevious(entities *entities) (Entity, bool, error) {

	sourceCodes, err := e.codeNode.ReadSources()
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourceCodes) {
	case 0:
		return e, true, nil
	case 1:
		parentCode := sourceCodes[0]

		bitIdentifier, characterIdentifier, err := parentCode.GetBitAndCharacter()
		if err != nil {
			return Entity{}, false, nil
		}

		return entities.create(bitIdentifier, parentCode.Identifier(), characterIdentifier), false, nil
	default:
		return Entity{}, false, fmt.Errorf("too many sources in code tree")
	}
}

func (e Entity) BitValue() (bool, error) {

	return e.bitNode.BitValue()
}

func (e Entity) String() string {
	return fmt.Sprintf("Code: %s %s %s\n", e.bitNode, e.codeNode, e.characterNode)
}
