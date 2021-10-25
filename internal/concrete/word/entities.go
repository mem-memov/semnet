package word

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/concrete/character"
	"github.com/mem-memov/semnet/internal/concrete/class"
	node2 "github.com/mem-memov/semnet/internal/concrete/word/node"
)

type entities struct {
	classes    *node2.Classes
	characters *node2.Characters
	words      *node2.Words
	phrases    *node2.Phrases
}

func newEntities(storage storage, classRepository *class.Repository, characterRepository *character.Repository) *entities {
	return &entities{
		classes:    node2.NewClasses(storage, classRepository),
		characters: node2.NewCharacters(storage, characterRepository),
		phrases:    node2.NewPhrases(storage),
		words:      node2.NewWords(storage),
	}
}

func (e *entities) create(classIdentifier uint, characterIdentifier uint, wordIdentifier uint, phraseIdentifier uint) Entity {
	return newEntity(
		e.classes.Create(classIdentifier),
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
	)
}

func (e *entities) createAndAddClass(characterIdentifier uint, wordIdentifier uint, phraseIdentifier uint) (Entity, error) {

	classNode, err := e.classes.CreateNew()
	if err != nil {
		return Entity{}, err
	}

	return newEntity(
		classNode,
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
	), nil
}

func (e *entities) createWithPhrase(phraseIdentifier uint) (Entity, error) {

	phraseNode := e.phrases.Create(phraseIdentifier)

	wordIdentifier, err := phraseNode.GetWord()
	if err != nil {
		return Entity{}, nil
	}

	wordNode := e.words.Create(wordIdentifier)

	classIdentifier, characterIdentifier, phraseIdentifierOfWord, err := wordNode.GetClassAndCharacterAndPhrase()
	if err != nil {
		return Entity{}, nil
	}

	if phraseIdentifier != phraseIdentifierOfWord {
		return Entity{}, fmt.Errorf("word has incorrect reference to phrase in word layer at word %d", wordIdentifier)
	}

	return e.create(classIdentifier, characterIdentifier, wordIdentifier, phraseIdentifier), nil
}
