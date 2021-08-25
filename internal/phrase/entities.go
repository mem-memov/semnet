package phrase

import (
	"github.com/mem-memov/semnet/internal/phrase/node"
	"github.com/mem-memov/semnet/internal/word"
)

type entities struct {
	words   *node.Words
	phrases *node.Phrases
	details *node.Details
}

func newEntities(storage storage, wordRepository *word.Repository) *entities {
	return &entities{
		words:   node.NewWords(storage, wordRepository),
		phrases: node.NewPhrases(storage),
		details: node.NewDetails(storage),
	}
}

func (e *entities) create(wordIdentifier uint, phraseIdentifier uint, detailIdentifier uint) Entity {
	return newEntity(
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
		e.details.Create(detailIdentifier),
	)
}
