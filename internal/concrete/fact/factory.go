package fact

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractFactClass "github.com/mem-memov/semnet/internal/abstract/fact/class"
	abstractFactPosition "github.com/mem-memov/semnet/internal/abstract/fact/position"
	abstractFactRemark "github.com/mem-memov/semnet/internal/abstract/fact/remark"
	abstractFactStory "github.com/mem-memov/semnet/internal/abstract/fact/story"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	"github.com/mem-memov/semnet/internal/concrete/fact/class"
	"github.com/mem-memov/semnet/internal/concrete/fact/position"
	"github.com/mem-memov/semnet/internal/concrete/fact/remark"
	"github.com/mem-memov/semnet/internal/concrete/fact/story"
)

type Factory struct {
	classFactory abstractFactClass.Factory
	remarkFactory abstractFactRemark.Factory
	positionFactory abstractFactPosition.Factory
	storyFactory abstractFactStory.Factory
}

var _ abstractFact.Factory = &Factory{}

func newFactory(storage abstract.Storage, classRepository abstractClass.Repository) *Factory {
	return &Factory{
		classFactory: class.NewFactory(classRepository),
		remarkFactory: remark.NewFactory(storage),
		positionFactory: position.NewFactory(storage),
		storyFactory: story.NewFactory(storage),
	}
}

func (f *Factory) CreateNewEntity(remarkEntity abstractRemark.Entity) (abstractFact.Entity, error) {
	classNode, err := f.classFactory.CreateNewNode()
	if err != nil {
		return nil, err
	}

	remarkNode, err := f.remarkFactory.CreateNewNode(remarkEntity, classNode)
	if err != nil {
		return nil, err
	}

	positionNode, err := f.positionFactory.CreateNewNode(remarkNode)
	if err != nil {
		return nil, err
	}

	storyNode, err := f.storyFactory.CreateNewNode(positionNode)
	if err != nil {
		return nil, err
	}

	err = remarkNode.GetMarked(remarkEntity)
	if err != nil {
		return nil, err
	}

	entity := newEntity(classNode, remarkNode, positionNode, storyNode)

	err = entity.GetMarked(remarkEntity)

	return entity, nil
}
