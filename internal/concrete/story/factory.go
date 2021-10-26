package story

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractFact "github.com/mem-memov/semnet/internal/abstract/fact"
	abstractStory "github.com/mem-memov/semnet/internal/abstract/story"
	abstractStoryClass "github.com/mem-memov/semnet/internal/abstract/story/class"
	abstractStoryFact "github.com/mem-memov/semnet/internal/abstract/story/fact"
	abstractStoryUser "github.com/mem-memov/semnet/internal/abstract/story/user"
	"github.com/mem-memov/semnet/internal/concrete/story/class"
	"github.com/mem-memov/semnet/internal/concrete/story/fact"
	"github.com/mem-memov/semnet/internal/concrete/story/user"
)

type factory struct {
	classFactory abstractStoryClass.Factory
	factFactory abstractStoryFact.Factory
	userFactory   abstractStoryUser.Factory
}

var _ abstractStory.Factory = &factory{}

func newFactory(storage abstract.Storage, classRepository abstractClass.Repository) *factory {
	return &factory{
		classFactory: class.NewFactory(classRepository),
		factFactory: fact.NewFactory(storage),
		userFactory:   user.NewFactory(storage),
	}
}

func (f *factory) CreateNewEntity(factEntity abstractFact.Entity) (abstractStory.Entity, error) {
	classNode, err := f.classFactory.CreateNewNode()
	if err != nil {
		return nil, err
	}

	factNode, err := f.factFactory.CreateNewNode(classNode)
	if err != nil {
		return nil, err
	}

	userNode, err := f.userFactory.CreateNewNode(factNode)
	if err != nil {
		return nil, err
	}

	entity := newEntity(classNode, factNode, userNode)

	err = entity.GetMarked(factEntity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
