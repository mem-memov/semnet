package remark

import "github.com/mem-memov/semnet/internal/detail"

type Repository struct {
	detailRepository detail.Repository
}

func NewRepository(detailRepository detail.Repository) *Repository {
	return &Repository{
		detailRepository: detailRepository,
	}
}

func (r *Repository) StartStory(object string, property string) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return Entity{}, err
	}

	entity, err := r.star.provideBeam(detailEntity)
	if err != nil {
		return Entity{}, err
	}

	return entity, nil
}

func (r *Repository) AppendFactToStory(previousIdentifier uint, object string, property string) (Entity, error) {

}

func (r *Repository) AddRemarkToFact(objectIdentifier uint, property string) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return Entity{}, err
	}

	return Entity{}, nil
}
