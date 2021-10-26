package fact

import (
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStoryFact "github.com/mem-memov/semnet/internal/abstract/story/fact"
)

type Node struct {
	identifier uint
}

var _ abstractStoryFact.Node = Node{}

func newNode(identifier uint) Node {
	return Node{
		identifier: identifier,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}

func (n Node) GetMarked(factEntity abstractFact.Entity) error {

	return factEntity.Mark(n.identifier)
}
