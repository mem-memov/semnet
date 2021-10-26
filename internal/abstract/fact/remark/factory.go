package remark

import (
	"github.com/mem-memov/semnet/internal/abstract/fact/class"
	"github.com/mem-memov/semnet/internal/abstract/remark"
)

type Factory interface {
	CreateNewNode(
		remarkEntity remark.Entity,
		classNode class.Node,
	) (Node, error)
}
