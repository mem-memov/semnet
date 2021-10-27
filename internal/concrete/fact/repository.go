package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Repository struct {
	storage abstract.Storage
	classRepository abstractClass.Repository
	storyRepository abstractStory.Repository
}

var _ abstractFact.Repository = &Repository{}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	storyRepository abstractStory.Repository,
) *Repository {
	return &Repository{
		storage: storage,
		classRepository: classRepository,
		storyRepository: storyRepository,
	}
}

func (r *Repository) CreateFirstUserStoryFact() (abstractFact.Entity, error) {

	fact, err :=  createEntity(r.storage)
	if err != nil {
		return nil, err
	}

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	err = fact.PointToClass(class)
	if err != nil {
		return nil, err
	}

	story, err := r.storyRepository.CreateFirstUserStory()
	if err != nil {
		return nil, err
	}

	err = fact.PointToStory(story)
	if err != nil {
		return nil, err
	}

	err = story.PointToFact(fact)
	if err != nil {
		return nil, err
	}

	return fact, nil
}
