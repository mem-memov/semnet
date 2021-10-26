package remark

import "github.com/mem-memov/semnet/internal/abstract/remark"

type Node interface {
	GetIdentifier() uint
	GetMarked(remarkEntity remark.Entity) error
}
