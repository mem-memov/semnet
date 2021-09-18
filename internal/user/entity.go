package user

import (
	"github.com/mem-memov/semnet/internal/story"
	"github.com/mem-memov/semnet/internal/user/node"
)

type Entity struct {
	classNode node.Class
	storyNode node.Story
}

func newEntity(classNode node.Class, storyNode node.Story) Entity {
	return Entity{
		classNode: classNode,
		storyNode: storyNode,
	}
}

func (e Entity) IdentifierForStory() uint {

	return e.storyNode.Identifier()
}

func (e Entity) AddStory() (story.Entity, error) {

	return story.Entity{}, nil
}

func (e Entity) GetStory(storyIdentifier uint) (story.Entity, error) {

	return story.Entity{}, nil
}
