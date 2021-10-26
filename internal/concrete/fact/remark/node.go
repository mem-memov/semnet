package remark

import (
	abstractFactRemark "github.com/mem-memov/semnet/internal/abstract/fact/remark"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Node struct {
	identifier uint
}

var _ abstractFactRemark.Node = Node{}

func newNode(identifier uint) Node {
	return Node{
		identifier: identifier,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}

func (n Node) GetMarked(remarkEntity abstractRemark.Entity) error {

	return remarkEntity.Mark(n.identifier)
}
