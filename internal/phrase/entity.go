package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/phrase/node"
)

type Entity struct {
	classNode  node.Class
	wordNode   node.Word
	phraseNode node.Phrase
	detailNode node.Detail
}

func newEntity(classNode node.Class, wordNode node.Word, phraseNode node.Phrase, detailNode node.Detail) Entity {
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

func (e Entity) ProvideDetailTarget(another Entity) (uint, error) {

	return e.detailNode.ProvideDetailTarget(another.detailNode)
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.detailNode.Mark(sourceIdentifier)
}

func (e Entity) provideNext(wordValue string, entities *entities) (Entity, error) {

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

		entity := entities.create(classIdentifier, wordIdentifier, targetPhrase.Identifier(), detailIdentifier)

		hasWordValue, err := entity.hasWordValue(wordValue)
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

func (e Entity) hasWordValue(wordValue string) (bool, error) {

	hasWordValue, err := e.wordNode.HasWordValue(wordValue)
	if err != nil {
		return false, err
	}

	return hasWordValue, nil
}

func (e Entity) wordValue() (string, error) {

	return e.wordNode.WordValue()
}

func (e Entity) findPrevious(entities *entities) (Entity, bool, error) {

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
			return Entity{}, false, nil
		}

		return entities.create(classIdentifier, wordIdentifier, parentPhrase.Identifier(), detailIdentifier), false, nil
	default:
		return Entity{}, false, fmt.Errorf("too many sources in phrase tree")
	}
}
