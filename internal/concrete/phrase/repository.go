package phrase

import (
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"github.com/mem-memov/semnet/internal/concrete/word"
)

type Repository struct {
	entities       *entities
	wordRepository *word.Repository
	tree           abstractPhrase.Tree
	paths          *paths
}

var _ abstractPhrase.Repository = &Repository{}

func NewRepository(storage storage, classRepository *class.Repository, wordRepository *word.Repository) *Repository {
	entities := newEntities(storage, classRepository, wordRepository)

	return &Repository{
		entities:       entities,
		wordRepository: wordRepository,
		tree:           newTree(storage, entities),
		paths:          newPaths(),
	}
}

func (r *Repository) Provide(words string) (abstractPhrase.Entity, error) {

	path, err := r.paths.collect(words)
	if err != nil {
		return Entity{}, err
	}

	firstWord, err := r.wordRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.tree.ProvideRoot(firstWord)
	if err != nil {
		return Entity{}, err
	}

	for _, wordValue := range path[1:] {

		entity, err = entity.ProvideNext(wordValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}

func (r *Repository) Extract(entity abstractPhrase.Entity) (string, error) {

	wordValue, err := entity.WordValue()
	if err != nil {
		return "", err
	}

	path := r.paths.create(wordValue)

	for {
		var isRoot bool
		entity, isRoot, err = entity.FindPrevious(r.entities)

		if isRoot {
			break
		}

		wordValue, err = entity.WordValue()
		if err != nil {
			return "", err
		}

		path = append(path, wordValue)
	}

	return path.reverse().toString(), nil
}

func (r *Repository) Fetch(detailIdentifier uint) (abstractPhrase.Entity, error) {

	return r.entities.CreateWithDetail(detailIdentifier)
}
