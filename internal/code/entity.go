package code

import (
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/code/node"
)

type Entity struct {
	bitNode        node.Bit
	codeNode       node.Code
	characterNode  node.Character
	bitRepository  bit.Repository
	codeRepository Repository
	storage        storage
}

func newEntity(bitNode node.Bit, codeNode node.Code, characterNode node.Character) Entity {
	return Entity{
		bitNode:       bitNode,
		codeNode:      codeNode,
		characterNode: characterNode,
	}
}

func (e Entity) createNext(bit bool) (Entity, error) {

	codeTargets, err := e.codeNode.readTargets()
	if err != nil {
		return Entity{}, err
	}

}
