package phrase

import (
	"fmt"
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

func (e *entities) createWithDetail(detailIdentifier uint) (Entity, error) {

	detailNode := e.details.Create(detailIdentifier)

	phraseIdentifier, err := detailNode.GetPhrase()
	if err != nil {
		return Entity{}, nil
	}

	phraseNode := e.phrases.Create(phraseIdentifier)

	wordIdentifier, detailIdentifierOfWord, err := phraseNode.GetWordAndDetail()
	if err != nil {
		return Entity{}, nil
	}

	if detailIdentifier != detailIdentifierOfWord {
		return Entity{}, fmt.Errorf("word has incorrect reference to detail in phrase layer at phrase %d", phraseIdentifier)
	}

	return e.create(wordIdentifier, phraseIdentifier, detailIdentifier), nil
}
