package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Repository struct {
	factStorage     abstractFact.Storage
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
		factStorage:     NewStorage(storage),
		classRepository: classRepository,
		storyRepository: storyRepository,
	}
}

func (r *Repository) CreateFirstUserStoryFact() (abstractFact.Aggregate, error) {

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateFact()
	if err != nil {
		return nil, err
	}

	fact, err := r.factStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	story, err := r.storyRepository.CreateFirstUserStory()
	if err != nil {
		return nil, err
	}

	err = fact.PointToStory(story.GetFact())
	if err != nil {
		return nil, err
	}

	err = story.PointToFact(fact.GetStory())
	if err != nil {
		return nil, err
	}

	return Aggregate{
		fact:            fact,
		factStorage:     r.factStorage,
		storyRepository: r.storyRepository,
	}, nil
}

func (r *Repository) FetchByRemark(remarkIdentifier uint) (abstractFact.Aggregate, error) {

	fact, err := r.factStorage.ReadEntityByRemark(remarkIdentifier)
	if err != nil {
		return nil, err
	}

	return Aggregate{
		fact:            fact,
		factStorage:     r.factStorage,
		storyRepository: r.storyRepository,
	}, nil
}

func (r *Repository) CreateNextFact(remarkIdentifier uint) (abstractFact.Aggregate, error) {

	fact, err := r.factStorage.ReadEntityByRemark(remarkIdentifier)
	if err != nil {
		return nil, err
	}

	story, err := r.storyRepository.FetchByFact(fact.GetStory())
	if err != nil {
		return nil, err
	}

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	classIdentifier, err := class.CreateFact()
	if err != nil {
		return nil, err
	}

	nextFact, err := r.factStorage.CreateEntity(classIdentifier)
	if err != nil {
		return nil, err
	}

	err = nextFact.PointToStory(story.GetFact())
	if err != nil {
		return nil, err
	}

	err = fact.PointToPosition(nextFact.GetPosition())
	if err != nil {
		return nil, err
	}

	return Aggregate{
		fact:            nextFact,
		factStorage:     r.factStorage,
		storyRepository: r.storyRepository,
	}, nil
}
