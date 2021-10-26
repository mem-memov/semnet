package fact

import "github.com/mem-memov/semnet/internal/abstract/remark"

type Factory interface {
	CreateNewEntity(remarkEntity remark.Entity) (Entity, error)
}
