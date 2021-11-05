package phrase

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Element struct {
	phrase        abstractPhrase.Entity
	phraseStorage abstractPhrase.Storage
	phraseFactory   abstractPhrase.Factory
	classRepository abstractClass.Repository
	wordRepository  abstractWord.Repository
}

func (e Element) GetNextElement(wordValue string) (abstractPhrase.Element, error) {

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
		targetWord, err := e.wordRepository.Fetch(targetPhrase.GetWord())
		if err != nil {
			return nil, err
		}

		targetValue, err := e.wordRepository.Extract(targetWord)
		if err != nil {
			return nil, err
		}

		if wordValue == targetValue {
			return Element{
				phrase: targetPhrase,
				phraseStorage: e.phraseStorage,
				phraseFactory: e.phraseFactory,
				classRepository: e.classRepository,
				wordRepository: e.wordRepository,
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

	newPhrase, err := e.phraseFactory.ProvideEntity(class, word)
	if err != nil {
		return nil, err
	}

	err = e.phrase.PointToPhrase(newPhrase.GetPhrase())
	if err != nil {
		return nil, err
	}

	return Element{
		phrase: newPhrase,
		phraseStorage: e.phraseStorage,
		phraseFactory: e.phraseFactory,
		classRepository: e.classRepository,
		wordRepository: e.wordRepository,
	}, nil
}