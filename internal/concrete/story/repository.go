package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Repository struct {
	factory abstractStory.Factory
}

var _ abstractStory.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository) *Repository {

	return &Repository{
		factory: newFactory(storage, classRepository),
	}
}

func (r *Repository) CreateNewEntity(factEntity abstractFact.Entity) (abstractStory.Entity, error) {

	return r.factory.CreateNewEntity(factEntity)
}
