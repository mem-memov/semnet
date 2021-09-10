package bit

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/class"
)

type Repository struct {
	storage storage
	layer   *layer
}

func NewRepository(storage storage, classRepository *class.Repository) *Repository {
	return &Repository{
		storage: storage,
		layer:   newLayer(storage, classRepository),
	}
}

func (r *Repository) Provide(value bool) (Entity, error) {

	zeroIdentifier, oneIdentifier, err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	if value {
		return newEntity(oneIdentifier, r.storage), nil
	} else {
		return newEntity(zeroIdentifier, r.storage), nil
	}
}

func (r *Repository) Fetch(identifier uint) (Entity, error) {

	zeroIdentifier, oneIdentifier, err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	if identifier != zeroIdentifier && identifier != oneIdentifier {
		return Entity{}, fmt.Errorf("wrong identifier in bit layer %d", identifier)
	}

	return newEntity(identifier, r.storage), nil
}
