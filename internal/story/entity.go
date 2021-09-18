package story

import (
	"github.com/mem-memov/semnet/internal/fact"
	"github.com/mem-memov/semnet/internal/story/node"
)

type Entity struct {
	factNode  node.Fact
	topicNode node.Topic
	storyNode node.Story
	userNode  node.User
}

func newEntity(factNode node.Fact, topicNode node.Topic, storyNode node.Story, userNode node.User) Entity {
	return Entity{
		factNode:  factNode,
		topicNode: topicNode,
		storyNode: storyNode,
		userNode:  userNode,
	}
}

func (e Entity) IdentifierForFact() uint {
	return e.factNode.Identifier()
}

func (e Entity) AddFact() (fact.Entity, error) {

	return fact.Entity{}, nil
}

func (e Entity) GetFact(factIdentifier uint) (fact.Entity, error) {

	return fact.Entity{}, nil
}
