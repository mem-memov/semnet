package word

import (
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Factory struct {
	wordStorage abstractWord.Storage
}

var _ abstractWord.Factory = &Factory{}

func NewFactory(wordStorage abstractWord.Storage) *Factory {

	return &Factory{
		wordStorage: wordStorage,
	}
}

func (f *Factory) ProvideFirstEntity(
	classEntity abstractClass.Entity,
	characterEntity abstractCharacter.Entity,
) (abstractWord.Entity, error) {

	hasCharacterSources, err := characterEntity.IsBeginningOfWords()
	if err != nil {
		return Entity{}, err
	}

	if !hasCharacterSources {
		return f.wordStorage.CreateEntity(classEntity, characterEntity)
	}

	character, err := characterEntity.ProvideSingleTarget()
	if err != nil {
		return nil, err
	}

	return f.wordStorage.ReadEntityByCharacter(character)
}
