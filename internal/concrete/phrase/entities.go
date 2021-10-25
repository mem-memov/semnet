package phrase

import (
	"fmt"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"github.com/mem-memov/semnet/internal/concrete/phrase/node"
	"github.com/mem-memov/semnet/internal/concrete/word"
)

type entities struct {
	classes abstractNode.Classes
	words   abstractNode.Words
	phrases abstractNode.Phrases
	details abstractNode.Details
}

var _ abstractPhrase.Entities = &entities{}

func newEntities(storage storage, classRepository *class.Repository, wordRepository *word.Repository) *entities {
	return &entities{
		classes: node.NewClasses(storage, classRepository),
		words:   node.NewWords(storage, wordRepository),
		phrases: node.NewPhrases(storage),
		details: node.NewDetails(storage),
	}
}

func (e *entities) Create(
	classIdentifier uint,
	wordIdentifier uint,
	phraseIdentifier uint,
	detailIdentifier uint,
) abstractPhrase.Entity {
	return newEntity(
		e.classes.Create(classIdentifier),
		e.words.Create(wordIdentifier),
		e.phrases.Create(phraseIdentifier),
		e.details.Create(detailIdentifier),
	)
}

func (e *entities) CreateAndAddClass(
	wordIdentifier uint,
	phraseIdentifier uint,
	detailIdentifier uint,
) (abstractPhrase.Entity, error) {

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

func (e *entities) CreateWithDetail(detailIdentifier uint) (abstractPhrase.Entity, error) {

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

	return e.Create(classIdentifier, wordIdentifier, phraseIdentifier, detailIdentifier), nil
}
