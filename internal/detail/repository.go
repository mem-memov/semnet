package detail

import (
	"github.com/mem-memov/semnet/internal/phrase"
)

type Repository struct {
	entities         *entities
	phraseRepository *phrase.Repository
	star             *star
}

func NewRepository(storage storage, phraseRepository *phrase.Repository) *Repository {
	entities := newEntities(storage, phraseRepository)

	return &Repository{
		entities:         entities,
		phraseRepository: phraseRepository,
		star:             newStar(storage, entities),
	}
}

func (r *Repository) Provide(object string, property string) (Entity, Entity, error) {

	objectPhrase, err := r.phraseRepository.Provide(object)
	if err != nil {
		return Entity{}, Entity{}, err
	}

	objectEntity, err := r.star.provideRoot(objectPhrase)
	if err != nil {
		return Entity{}, Entity{}, err
	}

	propertyEntity, err := objectEntity.provideNext(property, r.entities)
	if err != nil {
		return Entity{}, Entity{}, err
	}

	return objectEntity, propertyEntity, nil
}

func (r *Repository) Extract(entity Entity) (string, []string, error) {

	propertyValue, err := entity.phraseValue()
	if err != nil {
		return "", "", err
	}

	entity, err = entity.findPrevious(r.entities)


	path := r.paths.create(propertyValue)

	for {
		var isRoot bool
		entity, isRoot, err = entity.findPrevious(r.entities)

		if isRoot {
			break
		}

		propertyValue, err = entity.wordValue()
		if err != nil {
			return "", "", err
		}

		path = append(path, propertyValue)
	}

	return path.reverse().toString(), "", nil
}
