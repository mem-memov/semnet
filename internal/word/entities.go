package word

import (
	"fmt"
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

func (e *entities) createWithPhrase(phraseIdentifier uint) (Entity, error) {

	phraseNode := e.phrases.Create(phraseIdentifier)

	wordIdentifier, err := phraseNode.GetWord()
	if err != nil {
		return Entity{}, nil
	}

	wordNode := e.words.Create(wordIdentifier)

	characterIdentifier, phraseIdentifierOfWord, err := wordNode.GetCharacterAndPhrase()
	if err != nil {
		return Entity{}, nil
	}

	if phraseIdentifier != phraseIdentifierOfWord {
		return Entity{}, fmt.Errorf("word has incorrect reference to phrase in word layer at word %d", wordIdentifier)
	}

	return e.create(characterIdentifier, wordIdentifier, phraseIdentifier), nil
}
