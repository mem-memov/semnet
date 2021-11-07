package character

import (
	"github.com/mem-memov/semnet/internal/abstract/bit"
	"github.com/mem-memov/semnet/internal/abstract/class"
)

type Factory interface {
	ProvideFirstEntity(classEntity class.Entity, bitEntity bit.Entity) (Entity, error)
}
