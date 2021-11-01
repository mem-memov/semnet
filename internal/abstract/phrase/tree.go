package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract/word"
)

type Tree interface {
	ProvideRoot(wordEntity word.Entity) (Entity, error)
}
