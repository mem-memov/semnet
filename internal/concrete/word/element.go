package word

import (
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Element struct {
	word                abstractWord.Entity
	wordStorage         abstractWord.Storage
	wordFactory         abstractWord.Factory
	classRepository     abstractClass.Repository
	characterRepository abstractCharacter.Repository
}

var _ abstractWord.Element = Element{}

func (e Element) GetEntity() abstractWord.Entity {

	return e.word
}

func (e Element) ProvideNextElement(characterValue rune) (abstractWord.Element, error) {

	targetWordIdentifiers, err := e.word.GetTargetWords()
	if err != nil {
		return nil, err
	}

	targetWords := make([]abstractWord.Entity, 0, len(targetWordIdentifiers))
	for _, targetWordIdentifier := range targetWordIdentifiers {

		targetWord, err := e.wordStorage.ReadEntityByPhrase(targetWordIdentifier)
		if err != nil {
			return nil, err
		}

		targetWords = append(targetWords, targetWord)
	}

	// use existing
	for _, targetWord := range targetWords {
		targetCharacter, err := e.characterRepository.Fetch(targetWord.GetCharacter())
		if err != nil {
			return nil, err
		}

		targetValue, err := e.characterRepository.Extract(targetCharacter)
		if err != nil {
			return nil, err
		}

		if characterValue == targetValue {
			return Element{
				word:                targetWord,
				wordStorage:         e.wordStorage,
				wordFactory:         e.wordFactory,
				classRepository:     e.classRepository,
				characterRepository: e.characterRepository,
			}, nil
		}
	}

	// create new
	class, err := e.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	character, err := e.characterRepository.Provide(characterValue)
	if err != nil {
		return nil, err
	}

	character_ := character.(abstractCharacter.Entity) // TODO: remove

	newWord, err := e.wordFactory.ProvideFirstEntity(class, character_)
	if err != nil {
		return nil, err
	}

	err = e.word.PointToPhrase(newWord.GetWord())
	if err != nil {
		return nil, err
	}

	return Element{
		word:                newWord,
		wordStorage:         e.wordStorage,
		wordFactory:         e.wordFactory,
		classRepository:     e.classRepository,
		characterRepository: e.characterRepository,
	}, nil
}

func (e Element) ExtractCharacterValue() (rune, error) {

	targetCharacterIdentifier, err := e.word.GetTargetCharacter()
	if err != nil {
		return 0, err
	}

	word, err := e.characterRepository.Fetch(targetCharacterIdentifier)
	if err != nil {
		return 0, err
	}

	return e.characterRepository.Extract(word)
}

func (e Element) HasPreviousElement() (bool, error) {

	return e.word.HasSourceWord()
}

func (e Element) GetPreviousElement() (abstractWord.Element, error) {

	wordIdentifier, err := e.word.GetSourceWord()
	if err != nil {
		return nil, err
	}

	word, err := e.wordStorage.ReadEntityByWord(wordIdentifier)
	if err != nil {
		return nil, err
	}

	return Element{
		word:                word,
		wordStorage:         e.wordStorage,
		wordFactory:         e.wordFactory,
		classRepository:     e.classRepository,
		characterRepository: e.characterRepository,
	}, nil
}
