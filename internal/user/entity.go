package user

import (
	"github.com/mem-memov/semnet/internal/story"
	"github.com/mem-memov/semnet/internal/user/node"
)

type Entity struct {
	storyNode node.Story
	userNode  node.User
}

func newEntity(storyNode node.Story, userNode node.User) Entity {
	return Entity{
		storyNode: storyNode,
		userNode:  userNode,
	}
}

func (e Entity) AddStory() (story.Entity, error) {

	return story.Entity{}, nil
}

func (e Entity) GetStory(storyIdentifier uint) (story.Entity, error) {

	return story.Entity{}, nil
}
