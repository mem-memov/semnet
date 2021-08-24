package word

import (
	"github.com/mem-memov/semnet/internal/character"
)

type Repository struct {
	entities            *entities
	characterRepository *character.Repository
	tree                *tree
	paths               *paths
}

func NewRepository(storage storage, characterRepository *character.Repository) *Repository {
	entities := newEntities(storage, characterRepository)

	return &Repository{
		entities:            entities,
		characterRepository: characterRepository,
		tree:                newLayer(storage, entities),
		paths:               newPaths(),
	}
}

func (r *Repository) Provide(word string) (Entity, error) {

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

func (r *Repository) Extract(entity Entity) (string, error) {

	characterValue, err := entity.characterValue()
	if err != nil {
		return "", err
	}

	path := r.paths.create(characterValue)

	for {
		var isRoot bool
		entity, isRoot, err = entity.findPrevious(r.entities)

		if isRoot {
			break
		}

		characterValue, err = entity.characterValue()
		if err != nil {
			return "", err
		}

		path = append(path, characterValue)
	}

	return path.reverse().toString(), nil
}
