package remark

import "github.com/mem-memov/semnet/internal/abstract/detail"

type Chain interface {
	CreateFirstLink(detailEntity detail.Entity) (Entity, error)
}
