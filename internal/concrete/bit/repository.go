package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/concrete/class"
)

type Repository struct {
	storage storage
	layer   *layer
}

func NewRepository(storage storage, classRepository *class.Repository) *Repository {
	entities := newEntities(storage, classRepository)

	return &Repository{
		storage: storage,
		layer:   newLayer(storage, entities, classRepository),
	}
}

func (r *Repository) Provide(value bool) (Entity, error) {

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

func (r *Repository) Fetch(identifier uint) (Entity, error) {

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
