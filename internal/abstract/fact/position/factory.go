package position

import "github.com/mem-memov/semnet/internal/abstract/fact/remark"

type Factory interface {
	CreateNewNode(remarkNode remark.Node) (Node, error)
}
