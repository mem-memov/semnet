package user

import abstractStoryUser "github.com/mem-memov/semnet/internal/abstract/story/user"

type Node struct {
	identifier uint
}

var _ abstractStoryUser.Node = Node{}

func newNode(identifier uint) Node {
	return Node{
		identifier: identifier,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}
