package code

import (
	"github.com/mem-memov/semnet/internal/bit"
)

type Repository struct {
	entities      *entities
	storage       storage
	bitRepository *bit.Repository
	layer         *layer
	paths         *paths
}

func newRepository(storage storage, bitRepository *bit.Repository) *Repository {
	return &Repository{
		entities:      newEntities(storage, bitRepository),
		storage:       storage,
		bitRepository: bitRepository,
		layer:         newLayer(storage),
		paths:         newPaths(),
	}
}

func (r *Repository) Provide(integer int32) (Entity, error) {

	path, err := r.paths.collect(integer)
	if err != nil {
		return Entity{}, err
	}

	firstBit, err := r.bitRepository.Provide(path[0])
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.layer.provideRoot(firstBit)
	if err != nil {
		return Entity{}, err
	}

	for _, bitValue := range path[1:] {

		entity, err = entity.provideNext(bitValue, r.entities)
		if err != nil {
			return Entity{}, err
		}
	}

	return entity, nil
}
