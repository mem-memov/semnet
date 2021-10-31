package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Repository struct {
	storage         abstract.Storage
	classRepository abstractClass.Repository
}

var _ abstractStory.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository) *Repository {

	return &Repository{
		storage:         storage,
		classRepository: classRepository,
	}
}

func (r *Repository) CreateFirstUserStory() (abstractStory.Entity, error) {

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	story, err := createEntity(r.storage, class)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func (r *Repository) FetchByFact(factIdentifier uint) (abstractStory.Entity, error) {

	return readEntityByFact(r.storage, factIdentifier)
}
