package position

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractRemarkRemark "github.com/mem-memov/semnet/internal/abstract/remark/position"
)

type Node struct {
	identifier uint
	storage    abstract.Storage
}

var _ abstractRemarkRemark.Node = Node{}

func newNode(identifier uint, storage abstract.Storage) Node {
	return Node{
		identifier: identifier,
		storage:    storage,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}
