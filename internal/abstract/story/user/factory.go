package user

import (
	"github.com/mem-memov/semnet/internal/abstract/story/fact"
)

type Factory interface {
	CreateNewNode(factNode fact.Node) (Node, error)
}
