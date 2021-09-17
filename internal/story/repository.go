package story

import (
	"github.com/mem-memov/semnet/internal/class"
	"github.com/mem-memov/semnet/internal/fact"
	"github.com/mem-memov/semnet/internal/remark"
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

func (r *Repository) CreateStory(userIdentifier uint, remarks []remark.Entity) (Entity, error) {

	return Entity{}, nil
}

func (r *Repository) GetStory(storyIdentifier uint) (Entity, error) {

	return Entity{}, nil
}
