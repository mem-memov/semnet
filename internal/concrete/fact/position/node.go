package position

import abstractFactPosition "github.com/mem-memov/semnet/internal/abstract/fact/position"

type Node struct {
	identifier uint
}

var _ abstractFactPosition.Node = Node{}

func newNode(identifier uint) Node {
	return Node{
		identifier: identifier,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}