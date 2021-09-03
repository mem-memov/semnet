package story

import (
	"github.com/mem-memov/semnet/internal/fact"
	"github.com/mem-memov/semnet/internal/remark"
	"github.com/mem-memov/semnet/internal/user"
)

type Repository struct {
	storage storage
	factRepository *fact.Repository
}

func (r *Repository) CreateStory(remarkIdentifiers []uint, userIdentifier uint, object string, property string) (remark.Entity, error) {

	var storyEntity Entity

	r.factRepository.CreateFact(remarkIdentifiers, storyEntity.FactIdentifier(), object, property)

	return remark.Entity{}, nil
}

func (r *Repository) AddFactToStory(userEntity user.Entity, userStoryIdentifier uint, object string, property string) (remark.Entity, error) {

	var storyEntity Entity

	r.factRepository.AddFactToStory(storyEntity, object, property)

	return remark.Entity{}, nil
}