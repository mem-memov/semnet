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

func (r *Repository) CreateUser() (Entity, error) {

	return Entity{}, nil
}

func (r *Repository) GetUser(userIdentifier uint) (Entity, error) {

	return Entity{}, nil
}
