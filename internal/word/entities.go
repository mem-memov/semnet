package word

import (
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/word/node"
)

type entities struct {
	characters *node.Characters
	words      *node.Words
	phrases    *node.Phrases
}

func newEntities(storage storage, characterRepository *character.Repository) *entities {
	return &entities{
		characters: node.NewCharacters(storage, characterRepository),
		phrases:    node.NewPhrases(storage),
		words:      node.NewWords(storage),
	}
}

func (e *entities) create(characterIdentifier uint, wordIdentifier uint, phraseIdentifier uint) Entity {
	return newEntity(
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
	)
}
