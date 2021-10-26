package story

import "github.com/mem-memov/semnet/internal/abstract/fact/position"

type Factory interface {
	CreateNewNode(positionNode position.Node) (Node, error)
}
