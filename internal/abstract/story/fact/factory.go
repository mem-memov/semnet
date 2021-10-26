package fact

import (
	"github.com/mem-memov/semnet/internal/abstract/story/class"
)

type Factory interface {
	CreateNewNode(classNode class.Node) (Node, error)
}
