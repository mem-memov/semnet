package phrase

import (
	"fmt"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"
)

type Entity struct {
	classNode  abstractNode.Class
	wordNode   abstractNode.Word
	phraseNode abstractNode.Phrase
	detailNode abstractNode.Detail
}

var _ abstractPhrase.Entity = Entity{}

func newEntity(
	classNode abstractNode.Class,
	wordNode abstractNode.Word,
	phraseNode abstractNode.Phrase,
	detailNode abstractNode.Detail,
) Entity {
	return Entity{
		classNode:  classNode,
		wordNode:   wordNode,
		phraseNode: phraseNode,
		detailNode: detailNode,
	}
}

func (e Entity) DetailIdentifier() uint {
	return e.detailNode.Identifier()
}

func (e Entity) GetDetailNode() abstractNode.Detail {
	return e.detailNode
}

func (e Entity) ProvideDetailTarget(another abstractPhrase.Entity) (uint, error) {

	return e.detailNode.ProvideDetailTarget(another.GetDetailNode())
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.detailNode.Mark(sourceIdentifier)
}

func (e Entity) ProvideNext(wordValue string, entities abstractPhrase.Entities) (abstractPhrase.Entity, error) {

	targetPhrases, err := e.phraseNode.ReadTargets()
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetPhrase := range targetPhrases {
		classIdentifier, wordIdentifier, detailIdentifier, err := targetPhrase.GetClassAndWordAndDetail()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.Create(classIdentifier, wordIdentifier, targetPhrase.Identifier(), detailIdentifier)

		hasWordValue, err := entity.HasWordValue(wordValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasWordValue {
			return entity, nil
		}
	}

	// Provide new
	newClassNode, err := e.classNode.NewClass()
	if err != nil {
		return Entity{}, nil
	}

	newWordNode, err := e.wordNode.NewWord(wordValue)
	if err != nil {
		return Entity{}, nil
	}

	newPhraseNode, err := e.phraseNode.NewPhrase(newWordNode)
	if err != nil {
		return Entity{}, nil
	}

	newDetailNode, err := e.detailNode.NewDetail(newPhraseNode)
	if err != nil {
		return Entity{}, nil
	}

	return newEntity(newClassNode, newWordNode, newPhraseNode, newDetailNode), nil
}

func (e Entity) HasWordValue(wordValue string) (bool, error) {

	hasWordValue, err := e.wordNode.HasWordValue(wordValue)
	if err != nil {
		return false, err
	}

	return hasWordValue, nil
}

func (e Entity) WordValue() (string, error) {

	return e.wordNode.WordValue()
}

func (e Entity) FindPrevious(entities abstractPhrase.Entities) (abstractPhrase.Entity, bool, error) {

	sourcePhrases, err := e.phraseNode.ReadSources()
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourcePhrases) {
	case 0:
		return e, true, nil
	case 1:
		parentPhrase := sourcePhrases[0]

		classIdentifier, wordIdentifier, detailIdentifier, err := parentPhrase.GetClassAndWordAndDetail()
		if err != nil {
			return nil, false, nil
		}

		return entities.Create(classIdentifier, wordIdentifier, parentPhrase.Identifier(), detailIdentifier), false, nil
	default:
		return nil, false, fmt.Errorf("too many sources in phrase tree")
	}
}
