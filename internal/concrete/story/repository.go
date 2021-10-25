package story

import (
	"github.com/mem-memov/semnet/internal/concrete/class"
	"github.com/mem-memov/semnet/internal/concrete/fact"
)

type Repository struct {
	storage        storage
	factRepository *fact.Repository
}

func NewRepository(storage storage, classRepository *class.Repository, factRepository *fact.Repository) *Repository {
	return &Repository{
		factRepository: factRepository,
	}
}

func (r *Repository) CreateStory() (Story, error) {

	return nil, nil
}
