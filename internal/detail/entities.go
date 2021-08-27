package detail

import (
	"github.com/mem-memov/semnet/internal/detail/node"
	"github.com/mem-memov/semnet/internal/phrase"
)

type entities struct {
	phrases *node.Phrases
	remarks *node.Remarks
}

func newEntities(storage storage, phraseRepository *phrase.Repository) *entities {
	return &entities{
		phrases: node.NewPhrases(storage, phraseRepository),
		remarks: node.NewRemarks(storage),
	}
}

func (e *entities) create(phraseIdentifier uint, remarkIdentifier uint) Entity {
	return newEntity(
		e.phrases.Create(phraseIdentifier),
		e.remarks.Create(remarkIdentifier),
	)
}

func (e *entities) createWithRemark(remarkIdentifier uint) (Entity, error) {

	remarkNode := e.remarks.Create(remarkIdentifier)

	phraseIdentifier, err := remarkNode.GetPhrase()
	if err != nil {
		return Entity{}, nil
	}

	return e.create(phraseIdentifier, remarkIdentifier), nil
}
