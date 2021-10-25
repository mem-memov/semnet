package fact

import (
	"github.com/mem-memov/semnet/internal/concrete/class"
	remark2 "github.com/mem-memov/semnet/internal/concrete/remark"
	"github.com/mem-memov/semnet/internal/remark"
)

type Repository struct {
	storage          storage
	remarkRepository *remark2.Repository
}

func NewRepository(storage storage, classRepository *class.Repository, remarkRepository *remark2.Repository) *Repository {
	return &Repository{
		remarkRepository: remarkRepository,
	}
}

func (r *Repository) CreateFact(storyIdentifier uint, remarks []remark.Entity) (Entity, error) {

	return Entity{}, nil
}

func (r *Repository) GetRemarkFact(remark remark.Entity) (Entity, error) {

	return Entity{}, nil
}
