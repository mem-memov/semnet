package code

import (
	"github.com/mem-memov/semnet/internal/bit"
)

type Entity struct {
	bitNode        uint
	codeNode       uint
	characterNode  uint
	bitRepository  bit.Repository
	codeRepository Repository
	storage        storage
}

func newEntity(bitNode uint, codeNode uint, characterNode uint) Entity {
	return Entity{
		bitNode:       bitNode,
		codeNode:      codeNode,
		characterNode: characterNode,
	}
}
