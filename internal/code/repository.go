package code

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/code/node"
)

type Repository struct {
	storage       storage
	bitRepository *bit.Repository
	layer         *layer
	paths         *paths
}

func newRepository(storage storage, bitRepository *bit.Repository, paths *paths) *Repository {
	return &Repository{
		storage:       storage,
		bitRepository: bitRepository,
		layer:         newLayer(storage),
		paths:         paths,
	}
}

func (r *Repository) create(integer int32) (Entity, error) {

	path, err := r.paths.collect(integer)
	if err != nil {
		return Entity{}, err
	}

	bitEntity := path[0]

	bitTarget, err := bitEntity.CreateSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	bitNode := node.NewBit(bitTarget)

	entity, err := r.layer.createEntity(bitNode)
	if err != nil {
		return Entity{}, err
	}

	for i, bitEntity := range path[1:] {
		entity.createNext()
	}
}
