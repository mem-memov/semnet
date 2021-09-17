package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/class"
	"github.com/mem-memov/semnet/internal/phrase/node"
	"github.com/mem-memov/semnet/internal/word"
)

type entities struct {
	classes *node.Classes
	words   *node.Words
	phrases *node.Phrases
	details *node.Details
}

func newEntities(storage storage, classRepository *class.Repository, wordRepository *word.Repository) *entities {
	return &entities{
		classes: node.NewClasses(storage, classRepository),
		words:   node.NewWords(storage, wordRepository),
		phrases: node.NewPhrases(storage),
		details: node.NewDetails(storage),
	}
}

func (e *entities) create(classIdentifier uint, wordIdentifier uint, phraseIdentifier uint, detailIdentifier uint) Entity {
	return newEntity(
		e.classes.Create(classIdentifier),
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
		e.details.Create(detailIdentifier),
	)
}

func (e *entities) createAndAddClass(wordIdentifier uint, phraseIdentifier uint, detailIdentifier uint) (Entity, error) {

	classNode, err := e.classes.CreateNew()
	if err != nil {
		return Entity{}, err
	}

	return newEntity(
		classNode,
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
		e.details.Create(detailIdentifier),
	), nil
}

func (e *entities) createWithDetail(detailIdentifier uint) (Entity, error) {

	detailNode := e.details.Create(detailIdentifier)

	phraseIdentifier, err := detailNode.GetPhrase()
	if err != nil {
		return Entity{}, nil
	}

	phraseNode := e.phrases.Create(phraseIdentifier)

	classIdentifier, wordIdentifier, detailIdentifierOfWord, err := phraseNode.GetClassAndWordAndDetail()
	if err != nil {
		return Entity{}, nil
	}

	if detailIdentifier != detailIdentifierOfWord {
		return Entity{}, fmt.Errorf("word has incorrect reference to detail in phrase layer at phrase %d", phraseIdentifier)
	}

	return e.create(classIdentifier, wordIdentifier, phraseIdentifier, detailIdentifier), nil
}
