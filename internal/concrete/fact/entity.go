package fact

import (
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractFactClass "github.com/mem-memov/semnet/internal/abstract/fact/class"
	abstractFactPosition "github.com/mem-memov/semnet/internal/abstract/fact/position"
	abstractFactRemark "github.com/mem-memov/semnet/internal/abstract/fact/remark"
	abstractFactStory "github.com/mem-memov/semnet/internal/abstract/fact/story"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Entity struct {
	classNode  abstractFactClass.Node
	remarkNode abstractFactRemark.Node
	positionNode   abstractFactPosition.Node
	storyNode  abstractFactStory.Node
}

var _ abstractFact.Entity = Entity{}

func newEntity(
	classNode  abstractFactClass.Node,
	remarkNode abstractFactRemark.Node,
	positionNode   abstractFactPosition.Node,
	storyNode  abstractFactStory.Node,
) Entity {
	return Entity{
		classNode: classNode,
		remarkNode: remarkNode,
		positionNode: positionNode,
		storyNode: storyNode,
	}
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.storyNode.Mark(sourceIdentifier)
}

func (e Entity) GetMarked(remarkEntity abstractRemark.Entity) error {
	return e.remarkNode.GetMarked(remarkEntity)
}

