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

	return Entity{}, nil
}

func (r *Repository) AppendFact(previousIdentifier uint, object string, property string) (Entity, error) {

}

func (r *Repository) AddRemark(previousIdentifier uint, object string, property string) (Entity, error) {

	detailEntity, err := r.detailRepository.Provide(object, property)
	if err != nil {
		return Entity{}, err
	}

	return Entity{}, nil
}