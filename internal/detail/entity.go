package detail

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/detail/node"
)

type Entity struct {
	phraseNode node.Phrase
	detailNode node.Detail
	remarkNode node.Remark
}

func newEntity(phraseNode node.Phrase, detailNode node.Detail, remarkNode node.Remark) Entity {
	return Entity{
		phraseNode: phraseNode,
		detailNode: detailNode,
		remarkNode: remarkNode,
	}
}

func (e Entity) provideNext(phraseValue string, entities *entities) (Entity, error) {

	targetDetails, err := e.detailNode.ReadTargets()
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetDetail := range targetDetails {
		phraseIdentifier, remarkIdentifier, err := targetDetail.GetPhraseAndRemark()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.create(phraseIdentifier, targetDetail.Identifier(), remarkIdentifier)

		hasPhraseValue, err := entity.hasPhraseValue(phraseValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasPhraseValue {
			return entity, nil
		}
	}

	// Provide new
	newPhraseNode, err := e.phraseNode.NewPhrase(phraseValue)
	if err != nil {
		return Entity{}, nil
	}

	newDetailNode, err := e.detailNode.NewDetail(newPhraseNode)
	if err != nil {
		return Entity{}, nil
	}

	newRemarkNode, err := e.remarkNode.NewRemark(newDetailNode)
	if err != nil {
		return Entity{}, nil
	}

	return newEntity(newPhraseNode, newDetailNode, newRemarkNode), nil
}

func (e Entity) hasPhraseValue(phraseValue string) (bool, error) {

	hasPhraseValue, err := e.phraseNode.HasPhraseValue(phraseValue)
	if err != nil {
		return false, err
	}

	return hasPhraseValue, nil
}

func (e Entity) phraseValue() (string, error) {

	return e.phraseNode.PhraseValue()
}

func (e Entity) findPrevious(entities *entities) (Entity, bool, error) {

	sourceDetails, err := e.detailNode.ReadSources()
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourceDetails) {
	case 0:
		return e, true, nil
	case 1:
		parentPhrase := sourceDetails[0]

		wordIdentifier, detailIdentifier, err := parentPhrase.GetWordAndDetail()
		if err != nil {
			return Entity{}, false, nil
		}

		return entities.create(wordIdentifier, parentPhrase.Identifier(), detailIdentifier), false, nil
	default:
		return Entity{}, false, fmt.Errorf("too many sources in phrase tree")
	}
}
