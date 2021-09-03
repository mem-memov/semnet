package fact

import (
	"github.com/mem-memov/semnet/internal/remark"
	"github.com/mem-memov/semnet/internal/story"
)

type Repository struct {
	storage storage
	remarkRepository remark.Repository
}

func (r *Repository) CreateFact(remarkIdentifiers []uint, storyIdentifier uint, object string, property string) (remark.Entity, error) {

	var factEntity Entity

	r.remarkRepository.CreateRemark(remarkIdentifiers, factEntity.RemarkIdentifier(), object, property)

	return remark.Entity{}, nil
}

func (r *Repository) AddRemarkToFact(storyEntity story.Entity, object string, property string) (remark.Entity, error) {

	var factEntity Entity

	r.remarkRepository.AddRemarkToFact(factEntity, object, property)

	return remark.Entity{}, nil
}