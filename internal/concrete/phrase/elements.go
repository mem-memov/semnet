package phrase

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Elements struct {
	phraseStorage   abstractPhrase.Storage
	phraseFactory   abstractPhrase.Factory
	classRepository abstractClass.Repository
	wordRepository  abstractWord.Repository
}

var _ abstractPhrase.Elements = &Elements{}

func NewElements(
	phraseStorage abstractPhrase.Storage,
	phraseFactory abstractPhrase.Factory,
	classRepository abstractClass.Repository,
	wordRepository abstractWord.Repository,
) *Elements {

	return &Elements{
		phraseStorage:   phraseStorage,
		phraseFactory:   phraseFactory,
		classRepository: classRepository,
		wordRepository:  wordRepository,
	}
}

func (e *Elements) ProvideFirstElement(wordValue string) (abstractPhrase.Element, error) {

	firstWord, err := e.wordRepository.Provide(wordValue)
	if err != nil {
		return nil, err
	}

	class, err := e.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	phrase, err := e.phraseFactory.ProvideFirstEntity(class, firstWord)
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

func (e *Elements) CreateLastElement(phrase abstractPhrase.Entity) abstractPhrase.Element {

	return Element{
		phrase:          phrase,
		phraseStorage:   e.phraseStorage,
		phraseFactory:   e.phraseFactory,
		classRepository: e.classRepository,
		wordRepository:  e.wordRepository,
	}
}
