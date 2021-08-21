package bit

import "fmt"

type Repository struct {
	storage storage
	layer   *layer
}

func NewRepository(storage storage) *Repository {
	return &Repository{
		storage: storage,
		layer:   newLayer(storage),
	}
}

func (r *Repository) Provide(value bool) (Entity, error) {

	err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	if value {
		return newEntity(bitOneNode, r.storage), nil
	} else {
		return newEntity(bitZeroNode, r.storage), nil
	}
}

func (r *Repository) Fetch(identifier uint) (Entity, error) {

	err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	switch identifier {
	case bitZeroNode:
		return newEntity(bitZeroNode, r.storage), nil
	case bitOneNode:
		return newEntity(bitOneNode, r.storage), nil
	default:
		return Entity{}, fmt.Errorf("wrong identifier in bit layer %d", identifier)
	}
}
