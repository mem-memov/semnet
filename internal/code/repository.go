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
	entities := newEntities(storage, bitRepository)

	return &Repository{
		entities:      entities,
		storage:       storage,
		bitRepository: bitRepository,
		layer:         newLayer(storage, entities),
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

func (r *Repository) Extract(entity Entity) (int32, error) {

	bitValue, err := entity.BitValue()
	if err != nil {
		return 0, err
	}

	path := r.paths.create(bitValue)

	for {
		var isRoot bool
		entity, isRoot, err = entity.findPrevious(r.entities)

		if isRoot {
			break
		}

		bitValue, err = entity.BitValue()
		if err != nil {
			return 0, err
		}

		path = append(path, bitValue)
	}

	var integer int32

	for _, bitValue := range path.reverse() {
		integer <<= 1
		if bitValue {
			integer += 1
		}
	}

	return integer, nil
}
