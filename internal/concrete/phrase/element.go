package phrase

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Element struct {
	phrase          abstractPhrase.Entity
	phraseStorage   abstractPhrase.Storage
	phraseFactory   abstractPhrase.Factory
	classRepository abstractClass.Repository
	wordRepository  abstractWord.Repository
}

var _ abstractPhrase.Element = Element{}

func (e Element) GetEntity() abstractPhrase.Entity {

	return e.phrase
}

func (e Element) ProvideNextElement(wordValue string) (abstractPhrase.Element, error) {

	targetPhraseIdentifiers, err := e.phrase.GetTargetPhrases()
	if err != nil {
		return nil, err
	}

	targetPhrases := make([]abstractPhrase.Entity, 0, len(targetPhraseIdentifiers))
	for _, targetPhraseIdentifier := range targetPhraseIdentifiers {

		targetPhrase, err := e.phraseStorage.ReadEntityByPhrase(targetPhraseIdentifier)
		if err != nil {
			return nil, err
		}

		targetPhrases = append(targetPhrases, targetPhrase)
	}

	// use existing
	for _, targetPhrase := range targetPhrases {
		targetWordIdentifier, err := targetPhrase.GetTargetWord()
		if err != nil {
			return nil, err
		}

		targetWord, err := e.wordRepository.Fetch(targetWordIdentifier)
		if err != nil {
			return nil, err
		}

		targetValue, err := targetWord.Extract()
		if err != nil {
			return nil, err
		}

		if wordValue == targetValue {
			return Element{
				phrase:          targetPhrase,
				phraseStorage:   e.phraseStorage,
				phraseFactory:   e.phraseFactory,
				classRepository: e.classRepository,
				wordRepository:  e.wordRepository,
			}, nil
		}
	}

	// create new
	class, err := e.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	word, err := e.wordRepository.Provide(wordValue)
	if err != nil {
		return nil, err
	}

	newPhrase, err := e.phraseFactory.ProvideFirstEntity(class, word)
	if err != nil {
		return nil, err
	}

	err = e.phrase.PointToPhrase(newPhrase.GetPhrase())
	if err != nil {
		return nil, err
	}

	return Element{
		phrase:          newPhrase,
		phraseStorage:   e.phraseStorage,
		phraseFactory:   e.phraseFactory,
		classRepository: e.classRepository,
		wordRepository:  e.wordRepository,
	}, nil
}

func (e Element) ExtractWordValue() (string, error) {

	targetWordIdentifier, err := e.phrase.GetTargetWord()
	if err != nil {
		return "", err
	}

	word, err := e.wordRepository.Fetch(targetWordIdentifier)
	if err != nil {
		return "", err
	}

	return word.Extract()
}

func (e Element) HasPreviousElement() (bool, error) {

	return e.phrase.HasSourcePhrase()
}

func (e Element) GetPreviousElement() (abstractPhrase.Element, error) {

	phraseIdentifier, err := e.phrase.GetSourcePhrase()
	if err != nil {
		return nil, err
	}

	phrase, err := e.phraseStorage.ReadEntityByPhrase(phraseIdentifier)
	if err != nil {
		return nil, err
	}

	return Element{
		phrase:          phrase,
		phraseStorage:   e.phraseStorage,
		phraseFactory:   e.phraseFactory,
		classRepository: e.classRepository,
		wordRepository:  e.wordRepository,
	}, nil
}
