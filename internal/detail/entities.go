package detail

import (
	"github.com/mem-memov/semnet/internal/detail/node"
	"github.com/mem-memov/semnet/internal/phrase"
)

type entities struct {
	phrases *node.Phrases
	details *node.Details
	remarks *node.Remarks
}

func newEntities(storage storage, phraseRepository *phrase.Repository) *entities {
	return &entities{
		phrases: node.NewPhrases(storage, phraseRepository),
		details: node.NewDetails(storage),
		remarks: node.NewRemarks(storage),
	}
}

func (e *entities) create(phraseIdentifier uint, detailIdentifier uint, remarkIdentifier uint) Entity {
	return newEntity(
		e.phrases.Create(phraseIdentifier),
		e.details.Create(detailIdentifier),
		e.remarks.Create(remarkIdentifier),
	)
}
