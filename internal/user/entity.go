package user

import (
	"github.com/mem-memov/semnet/internal/remark"
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

func (e Entity) AddStoryToUser(remarkIdentifiers []uint, object string, property string) (remark.Entity, error) {

	remarkEntity, err := e.storyNode.CreateStory(remarkIdentifiers, object, property)
	if err != nil {
		return remark.Entity{}, err
	}

	return remarkEntity, nil
}

func (e Entity) AddFactToStory(remarkIdentifiers []uint, object string, property string) (remark.Entity, error) {

	return remark.Entity{}, nil
}

func (e Entity) AddRemarkToFact(remarkIdentifiers []uint, object string, property string) (remark.Entity, error) {

	return remark.Entity{}, nil
}
