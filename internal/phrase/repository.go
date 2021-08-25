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

func (r *Repository) Provide(words string) (Entity, error) {

	path, err := r.paths.collect(words)
	if err != nil {
		return Entity{}, err
	}

	firstWord, err := r.wordRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.tree.provideRoot(firstWord)
	if err != nil {
		return Entity{}, err
	}

	for _, wordValue := range path[1:] {

		entity, err = entity.provideNext(wordValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}

func (r *Repository) Extract(entity Entity) (string, error) {

	wordValue, err := entity.wordValue()
	if err != nil {
		return "", err
	}

	path := r.paths.create(wordValue)

	for {
		var isRoot bool
		entity, isRoot, err = entity.findPrevious(r.entities)

		if isRoot {
			break
		}

		wordValue, err = entity.wordValue()
		if err != nil {
			return "", err
		}

		path = append(path, wordValue)
	}

	return path.reverse().toString(), nil
}

func (r *Repository) Fetch(detailIdentifier uint) (Entity, error) {

	return r.entities.createWithDetail(detailIdentifier)
}
