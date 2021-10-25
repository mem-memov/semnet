package phrase

import "github.com/mem-memov/semnet/internal/concrete/word"

type Tree interface {
	ProvideRoot(wordEntity word.Entity) (Entity, error)
}
