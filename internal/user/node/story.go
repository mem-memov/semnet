package node

import (
	"github.com/mem-memov/semnet/internal/remark"
	"github.com/mem-memov/semnet/internal/story"
)

type Story struct {
	identifier      uint
	storage         storage
	storyRepository *story.Repository
}

func (s Story) CreateStory(remarkIdentifiers []uint, object string, property string) (remark.Entity, error) {

	return s.storyRepository.CreateStory(remarkIdentifiers, s.identifier, object, property)
}
