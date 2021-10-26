package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
)

type Repository struct {
	factory abstractFact.Factory
	storyRepository abstractStory.Repository
}

func NewRepository(
	storage abstract.Storage,
	classRepository abstractClass.Repository,
	storyRepository abstractStory.Repository,
) *Repository {
	return &Repository{
		factory: newFactory(storage, classRepository),
		storyRepository: storyRepository,
	}
}

func (r *Repository) CreateNewFact(remarkEntity abstractRemark.Entity) (abstractFact.Entity, error) {

	entity, err :=  r.factory.CreateNewEntity(remarkEntity)
	if err != nil {
		return nil, err
	}

	_, err = r.storyRepository.CreateNewEntity(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
