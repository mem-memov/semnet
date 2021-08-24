package word

import (
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
