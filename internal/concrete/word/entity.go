package word

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
	"github.com/mem-memov/semnet/internal/concrete/word/node"
)

type Entity struct {
	class     uint
	character uint
	word      uint
	phrase    uint
	storage   abstract.Storage
}

var _ abstractWord.Entity = Entity{}

func (e Entity) PointToPhrase(phrase uint) error {

	return e.storage.Connect(e.phrase, phrase)
}

func (e Entity) PhraseIdentifier() uint {

	return e.phrase
}

func (e Entity) Mark(sourceIdentifier uint) error {

	return e.storage.Connect(sourceIdentifier, e.phrase)
}

func (e Entity) ProvideSingleTarget() (uint, error) {

	targets, err := e.storage.ReadTargets(e.phrase)
	if err != nil {
		return 0, err
	}

	switch len(targets) {

	case 0:
		target, err := e.storage.Create()
		if err != nil {
			return 0, err
		}

		err = e.storage.Connect(e.phrase, target)
		if err != nil {
			return 0, err
		}

		return target, nil

	case 1:
		return targets[0], nil

	default:
		return 0, fmt.Errorf("cluster %d has too many targets: %d", e.phrase, len(targets))
	}
}

func (e Entity) HasSingleTargetOtherTargets() (bool, error) {

	target, err := e.ProvideSingleTarget()
	if err != nil {
		return false, err
	}

	backTargets, err := e.storage.ReadTargets(target)

	switch len(backTargets) {

	case 0:

		return false, nil

	case 1:

		if backTargets[0] != e.phrase {
			return false, fmt.Errorf("word not pointing to itself: %d", e.phrase)
		}

		return true, nil

	default:

		return false, fmt.Errorf("word not pointing to itself: %d", e.phrase)
	}
}

func (e Entity) ProvideNext(characterValue rune, entities *entities) (Entity, error) {

	targetWords, err := e.storage.ReadTargets(e.word)
	if err != nil {
		return Entity{}, nil
	}

	// search existing
	for _, targetWord := range targetWords {
		classIdentifier, characterIdentifier, phraseIdentifier, err := targetWord.GetClassAndCharacterAndPhrase()
		if err != nil {
			return Entity{}, nil
		}

		entity := entities.create(classIdentifier, characterIdentifier, targetWord.Identifier(), phraseIdentifier)

		hasBitValue, err := entity.HasCharacterValue(characterValue)
		if err != nil {
			return Entity{}, nil
		}

		if hasBitValue {
			return entity, nil
		}
	}

	// Provide new
	newClassNode, err := e.classNode.NewClass()
	if err != nil {
		return Entity{}, nil
	}

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

	return newEntity(newClassNode, newCharacterNode, newWordNode, newPhraseNode), nil
}

func (e Entity) HasCharacterValue(characterValue rune) (bool, error) {

	hasCharacterValue, err := e.characterNode.HasCharacterValue(characterValue)
	if err != nil {
		return false, err
	}

	return hasCharacterValue, nil
}

func (e Entity) CharacterValue() (rune, error) {

	return e.characterNode.CharacterValue()
}

func (e Entity) FindPrevious(entities *entities) (Entity, bool, error) {

	sourceWords, err := e.storage.ReadSources(e.word)
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourceWords) {
	case 0:
		return e, true, nil
	case 1:
		parentWord := sourceWords[0]

		class, character, phrase, err := parentWord.GetClassAndCharacterAndPhrase()
		if err != nil {
			return Entity{}, false, nil
		}

		return entities.create(class, character, parentWord, phrase), false, nil
	default:
		return Entity{}, false, fmt.Errorf("too many sources in word tree")
	}
}
