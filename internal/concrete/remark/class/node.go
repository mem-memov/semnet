package class

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractRemarkClass "github.com/mem-memov/semnet/internal/abstract/remark/class"
)

// Node structure wraps a class node identifier.
type Node struct {
	identifier uint
	repository abstractClass.Repository
	creator    abstractRemarkClass.IdentifierCreator
}

var _ abstractRemarkClass.Node = Node{}

func newNode(
	identifier uint,
	repository abstractClass.Repository,
	creator abstractRemarkClass.IdentifierCreator,
) abstractRemarkClass.Node {
	return Node{
		identifier: identifier,
		repository: repository,
		creator:    creator,
	}
}

func (n Node) GetIdentifier() uint {
	return n.identifier
}

// IsValid checks the single target of this node is common to all position entities.
func (n Node) IsValid() (bool, error) {
	classEntity, err := n.repository.ProvideEntity()
	if err != nil {
		return false, err
	}

	return classEntity.IsRemark(n.identifier)
}

// CreateNewNode takes data of the present node and passes it to a new node.
// All nodes of an cluster must have such ability so that the cluster can produce a new cluster building a chain of entities.
func (n Node) CreateNewNode() (abstractRemarkClass.Node, error) {
	newIdentifier, err := n.creator.CreateNewIdentifier()
	if err != nil {
		return Node{}, err
	}

	return newNode(newIdentifier, n.repository, n.creator), nil
}
