package story

import (
	node2 "github.com/mem-memov/semnet/internal/concrete/story/node"
	"github.com/mem-memov/semnet/internal/story/node"
)

type Entity struct {
	factNode  node2.Fact
	topicNode node.Topic
	storyNode node2.Story
}

func newEntity(factNode node2.Fact, topicNode node.Topic, storyNode node2.Story) Entity {
	return Entity{
		factNode:  factNode,
		topicNode: topicNode,
		storyNode: storyNode,
	}
}

func (e Entity) IdentifierForFact() uint {
	return e.factNode.Identifier()
}

func (e Entity) AddFact() (Fact, error) {

	return nil, nil
}
