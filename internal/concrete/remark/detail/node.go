package detail

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractCreator "github.com/mem-memov/semnet/internal/abstract/remark/detail"
)

type Node struct {
	identifier uint
	storage    abstract.Storage
	repository abstractDetail.Repository
	creator    abstractCreator.IdentifierCreator
}

func newNode(
	identifier uint,
	storage abstract.Storage,
	repository abstractDetail.Repository,
	creator abstractCreator.IdentifierCreator,
) Node {
	return Node{
		identifier:       identifier,
		storage:          storage,
		repository: repository,
		creator: creator,
	}
}

// GetObjectAndProperty provides actual text of the remark
func (n Node) GetObjectAndProperty() (string, string, error) {
	detailEntity, err := n.repository.Fetch(n.identifier)
	if err != nil {
		return "", "", err
	}

	return detailEntity.GetObjectAndProperty()
}

// CreateNewNode takes data of the present node and passes it to a new node.
// All nodes of a cluster must have such ability so that the cluster can produce a new cluster building a chain of entities.
func (n Node) CreateNewNode(object string, property string) (Node, error) {
	identifier, err := n.creator.CreateNewIdentifier(object, property )
	if err != nil {
		return Node{}, err
	}

	return newNode(identifier, n.storage, n.repository, n.creator), nil
}


