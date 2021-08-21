package code

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
)

type Repository struct {
	storage       storage
	bitRepository *bit.Repository
	layer         *layer
}

func newRepository(storage storage, bitRepository *bit.Repository) *Repository {
	return &Repository{
		storage:       storage,
		bitRepository: bitRepository,
		layer:         newLayer(storage),
	}
}

func (r *Repository) create(integer int32) (Entity, error) {

	bitNames := fmt.Sprintf("%b", integer)

	if len(bitNames) < 1 {
		return Entity{}, fmt.Errorf("no bitEntities in Entity: %d", integer)
	}

	bitEntities := make([]bit.Entity, len(bitNames))

	for i, bitName := range bitNames {
		if bitName != '0' && bitName != '1' {
			return Entity{}, fmt.Errorf("invalid bitEntity name: %c", bitName)
		}

		bitEntity, err := r.bitRepository.Create(bitName == '1')
		if err != nil {
			return Entity{}, err
		}

		bitEntities[i] = bitEntity
	}

	bitEntity := bitEntities[0]

	bitNode, err := bitEntity.GetSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.layer.createEntity(bitNode)
	if err != nil {
		return Entity{}, err
	}
}
