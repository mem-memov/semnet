package fact

import (
	"github.com/mem-memov/semnet/internal/class"
	"github.com/mem-memov/semnet/internal/remark"
)

type Repository struct {
	storage          storage
	remarkRepository *remark.Repository
}

func NewRepository(storage storage, classRepository *class.Repository, remarkRepository *remark.Repository) *Repository {
	return &Repository{
		remarkRepository: remarkRepository,
	}
}

func (r *Repository) CreateFact(storyIdentifier uint, remarks []remark.Entity) (Entity, error) {

	return Entity{}, nil
}

func (r *Repository) GetFact(identifier uint) (Entity, error) {

	return Entity{}, nil
}
