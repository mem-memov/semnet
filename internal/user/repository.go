package user

import (
	"github.com/mem-memov/semnet/internal/story"
)

type Repository struct {
	storage         storage
	storyRepository *story.Repository
}

func NewRepository(storage storage, storyRepository *story.Repository) *Repository {
	return &Repository{
		storage:         storage,
		storyRepository: storyRepository,
	}
}

func (r *Repository) Fetch(userIdentifier uint) (Entity, error) {

	return Entity{}, nil
}
