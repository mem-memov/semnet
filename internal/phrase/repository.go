package phrase

import (
	"github.com/mem-memov/semnet/internal/word"
)

type Repository struct {
	entities       *entities
	wordRepository *word.Repository
	tree           *tree
	paths          *paths
}

func NewRepository(storage storage, wordRepository *word.Repository) *Repository {
	entities := newEntities(storage, wordRepository)

	return &Repository{
		entities:       entities,
		wordRepository: wordRepository,
		tree:           newTree(storage, entities),
		paths:          newPaths(),
	}
}
