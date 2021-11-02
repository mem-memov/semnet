package word

import (
	"github.com/mem-memov/semnet/internal/concrete/character"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Repository struct {
	entities            *entities
	characterRepository *character.Repository
	tree                *tree
	paths               *paths
}

func NewRepository(storage storage, classRepository *class.Repository, characterRepository *character.Repository) *Repository {
	entities := newEntities(storage, classRepository, characterRepository)

	return &Repository{
		entities:            entities,
		characterRepository: characterRepository,
		tree:                newTree(storage, entities),
		paths:               newPaths(),
	}
}

func (r *Repository) Provide(word string) (interface{}, error) {

	path, err := r.paths.collect(word)
	if err != nil {
		return Entity{}, err
	}

	firstCharacter, err := r.characterRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.tree.provideRoot(firstCharacter)
	if err != nil {
		return Entity{}, err
	}

	for _, characterValue := range path[1:] {

		entity, err = entity.provideNext(characterValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}

func (r *Repository) Extract(entity interface{}) (string, error) {

	// TODO: remove after rafactoring
	entit := entity.(Entity)

	characterValue, err := entit.characterValue()
	if err != nil {
		return "", err
	}

	path := r.paths.create(characterValue)

	for {
		var isRoot bool
		entity, isRoot, err = entit.findPrevious(r.entities)

		if isRoot {
			break
		}

		characterValue, err = entit.characterValue()
		if err != nil {
			return "", err
		}

		path = append(path, characterValue)
	}

	return path.reverse().toString(), nil
}

func (r *Repository) Fetch(phraseIdentifier uint) (interface{}, error) {

	return r.entities.createWithPhrase(phraseIdentifier)
}
