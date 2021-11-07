package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Repository struct {
	storage abstract.Storage
	layer   *layer
}

var _ abstractBit.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository *class.Repository) *Repository {
	entities := newEntities(storage, classRepository)

	return &Repository{
		storage: storage,
		layer:   newLayer(storage, entities, classRepository),
	}
}

func (r *Repository) Provide(value bool) (abstractBit.Entity, error) {

	zeroEntity, oneEntity, err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	if value {
		return oneEntity, nil
	} else {
		return zeroEntity, nil
	}
}

func (r *Repository) Fetch(identifier uint) (abstractBit.Entity, error) {

	zeroEntity, oneEntity, err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	if identifier == zeroEntity.Identifier() {
		return zeroEntity, nil
	}

	if identifier == oneEntity.Identifier() {
		return oneEntity, nil
	}

	return Entity{}, fmt.Errorf("wrong identifier in bit layer %d", identifier)
}
