package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractFactStory "github.com/mem-memov/semnet/internal/abstract/fact/story"
)

type Node struct {
	identifier uint
	storage abstract.Storage
}

var _ abstractFactStory.Node = Node{}

func newNode(identifier uint, storage abstract.Storage) Node {
	return Node{
		identifier: identifier,
		storage: storage,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}

func (n Node) Mark(sourceIdentifier uint) error {
	return n.storage.Connect(sourceIdentifier, n.identifier)
}