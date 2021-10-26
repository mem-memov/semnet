package story

import (
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStoryClass "github.com/mem-memov/semnet/internal/abstract/story/class"
	abstractStoryFact "github.com/mem-memov/semnet/internal/abstract/story/fact"
	abstractStoryUser "github.com/mem-memov/semnet/internal/abstract/story/user"
)

type Entity struct {
	classNode  abstractStoryClass.Node
	factNode abstractStoryFact.Node
	userNode abstractStoryUser.Node
}

func newEntity(
	classNode  abstractStoryClass.Node,
	factNode abstractStoryFact.Node,
	userNode abstractStoryUser.Node,
) Entity {
	return Entity{
		classNode:  classNode,
		factNode: factNode,
		userNode: userNode,
	}
}

func (e Entity) GetMarked(factEntity abstractFact.Entity) error {

	return e.factNode.GetMarked(factEntity)
}


