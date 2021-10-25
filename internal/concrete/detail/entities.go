package detail

import (
	"github.com/mem-memov/semnet/internal/abstract/phrase"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"github.com/mem-memov/semnet/internal/concrete/detail/node"
)

type entities struct {
	classes *node.Classes
	phrases *node.Phrases
	remarks *node.Remarks
}

func newEntities(storage storage, classRepository *class.Repository, phraseRepository phrase.Repository) *entities {
	return &entities{
		classes: node.NewClasses(storage, classRepository),
		phrases: node.NewPhrases(storage, phraseRepository),
		remarks: node.NewRemarks(storage),
	}
}

func (e *entities) create(classIdentifier uint, phraseIdentifier uint, remarkIdentifier uint) Entity {
	return newEntity(
		e.classes.Create(classIdentifier),
		e.phrases.Create(phraseIdentifier),
		e.remarks.Create(remarkIdentifier),
	)
}

func (e *entities) createAndAddClass(phraseIdentifier uint, remarkIdentifier uint) (Entity, error) {

	classNode, err := e.classes.CreateNew()
	if err != nil {
		return Entity{}, err
	}

	return newEntity(
		classNode,
		e.phrases.Create(phraseIdentifier),
		e.remarks.Create(remarkIdentifier),
	), nil
}

func (e *entities) createWithRemark(remarkIdentifier uint) (Entity, error) {

	remarkNode := e.remarks.Create(remarkIdentifier)

	classIdentifier, phraseIdentifier, err := remarkNode.GetClassAndPhrase()
	if err != nil {
		return Entity{}, nil
	}

	return e.create(classIdentifier, phraseIdentifier, remarkIdentifier), nil
}
