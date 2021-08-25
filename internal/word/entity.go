package word

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/word/node"
)

type Entity struct {
	characterNode node.Character
	wordNode      node.Word
	phraseNode    node.Phrase
}

func newEntity(characterNode node.Character, wordNode node.Word, phraseNode node.Phrase) Entity {
	return Entity{
		characterNode: characterNode,
		wordNode:      wordNode,
		phraseNode:    phraseNode,
	}
}

func (e Entity) PhraseIdentifier() uint {
	return e.phraseNode.Identifier()
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.phraseNode.Mark(sourceIdentifier)
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	return e.phraseNode.ProvideSingleTarget()
}

func (e Entity) provideNext(characterValue rune, entities *entities) (Entity, error) {

	targetWords, err := e.wordNode.ReadTargets()
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetWord := range targetWords {
		characterIdentifier, phraseIdentifier, err := targetWord.GetCharacterAndPhrase()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.create(characterIdentifier, targetWord.Identifier(), phraseIdentifier)

		hasBitValue, err := entity.hasCharacterValue(characterValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasBitValue {
			return entity, nil
		}
	}

	// Provide new
	newCharacterNode, err := e.characterNode.NewCharacter(characterValue)
	if err != nil {
		return Entity{}, nil
	}

	newWordNode, err := e.wordNode.NewWord(newCharacterNode)
	if err != nil {
		return Entity{}, nil
	}

	newPhraseNode, err := e.phraseNode.NewPhrase(newWordNode)
	if err != nil {
		return Entity{}, nil
	}

	return newEntity(newCharacterNode, newWordNode, newPhraseNode), nil
}

func (e Entity) hasCharacterValue(characterValue rune) (bool, error) {

	hasCharacterValue, err := e.characterNode.HasCharacterValue(characterValue)
	if err != nil {
		return false, err
	}

	return hasCharacterValue, nil
}

func (e Entity) characterValue() (rune, error) {

	return e.characterNode.CharacterValue()
}

func (e Entity) findPrevious(entities *entities) (Entity, bool, error) {

	sourceWords, err := e.wordNode.ReadSources()
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourceWords) {
	case 0:
		return e, true, nil
	case 1:
		parentWord := sourceWords[0]

		characterIdentifier, phraseIdentifier, err := parentWord.GetCharacterAndPhrase()
		if err != nil {
			return Entity{}, false, nil
		}

		return entities.create(characterIdentifier, parentWord.Identifier(), phraseIdentifier), false, nil
	default:
		return Entity{}, false, fmt.Errorf("too many sources in word tree")
	}
}
