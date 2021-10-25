package bit

import (
	node2 "github.com/mem-memov/semnet/internal/concrete/bit/node"
)

type Entity struct {
	value         bool
	classNode     node2.Class
	characterNode node2.Character
}

func newEntity(value bool, classNode node2.Class, characterNode node2.Character) Entity {
	return Entity{
		value:         value,
		classNode:     classNode,
		characterNode: characterNode,
	}
}

func (e Entity) Identifier() uint {
	return e.characterNode.Identifier()
}

func (e Entity) Is(bit bool) bool {
	return e.value == bit
}

func (e Entity) Bit() bool {
	return e.value
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.characterNode.Mark(sourceIdentifier)
}

func (e Entity) ProvideSingleTarget() (uint, error) {
	return e.characterNode.ProvideSingleTarget()
}
