package word

import (
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Elements struct {
	wordStorage         abstractWord.Storage
	wordFactory         abstractWord.Factory
	classRepository     abstractClass.Repository
	characterRepository abstractCharacter.Repository
}

var _ abstractWord.Elements = &Elements{}

func NewElements(
	wordStorage abstractWord.Storage,
	wordFactory abstractWord.Factory,
	classRepository abstractClass.Repository,
	characterRepository abstractCharacter.Repository,
) *Elements {

	return &Elements{
		wordStorage:         wordStorage,
		wordFactory:         wordFactory,
		classRepository:     classRepository,
		characterRepository: characterRepository,
	}
}

func (e *Elements) ProvideFirstElement(characterValue rune) (abstractWord.Element, error) {

	firstCharacter, err := e.characterRepository.Provide(characterValue)
	if err != nil {
		return nil, err
	}

	class, err := e.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	word, err := e.wordFactory.ProvideHeadEntity(class, firstCharacter)
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

func (e *Elements) CreateLastElement(word abstractWord.Entity) abstractWord.Element {

	return Element{
		word:                word,
		wordStorage:         e.wordStorage,
		wordFactory:         e.wordFactory,
		classRepository:     e.classRepository,
		characterRepository: e.characterRepository,
	}
}
