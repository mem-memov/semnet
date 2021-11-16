package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Repository struct {
	storyStorage    abstractStory.Storage
	classRepository abstractClass.Repository
}

var _ abstractStory.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository) *Repository {

	return &Repository{
		storyStorage:    NewStorage(storage),
		classRepository: classRepository,
	}
}

func (r *Repository) CreateStory() (abstractStory.Entity, error) {

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateStory()
	if err != nil {
		return nil, err
	}

	story, err := r.storyStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func (r *Repository) FetchByFact(factIdentifier uint) (abstractStory.Entity, error) {

	return r.storyStorage.ReadEntityByFact(factIdentifier)
}

func (r *Repository) FetchByPosition(positionIdentifier uint) (abstractStory.Entity, error) {

	return r.storyStorage.ReadEntityByPosition(positionIdentifier)
}

func (r *Repository) FetchByTree(treeIdentifier uint) (abstractStory.Entity, error) {

	return r.storyStorage.ReadEntityByTree(treeIdentifier)
}
