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

func (f *Factory) ProvideHeadEntity(
	classEntity abstractClass.Entity,
	characterAggregate abstractCharacter.Aggregate,
) (abstractWord.Entity, error) {

	hasTargetWord, err := characterAggregate.HasTargetWord()
	if err != nil {
		return Entity{}, err
	}

	if !hasTargetWord {

		class, err := classEntity.CreateWord()
		if err != nil {
			return nil, err
		}

		wordEntity, err := f.wordStorage.CreateEntity(class)
		if err != nil {
			return nil, err
		}

		err = wordEntity.PointToCharacter(characterAggregate.GetWord())
		if err != nil {
			return nil, err
		}

		err = characterAggregate.PointToWord(wordEntity.GetCharacter())
		if err != nil {
			return nil, err
		}

		return wordEntity, nil
	}

	character, err := characterAggregate.GetTargetWord()
	if err != nil {
		return nil, err
	}

	return f.wordStorage.ReadEntityByCharacter(character)
}

func (f *Factory) CreateTailEntity(
	classEntity abstractClass.Entity,
	characterAggregate abstractCharacter.Aggregate,
	previousWordEntity abstractWord.Entity,
) (abstractWord.Entity, error) {

	class, err := classEntity.CreateWord()
	if err != nil {
		return nil, err
	}

	newWordEntity, err := f.wordStorage.CreateEntity(class)
	if err != nil {
		return nil, err
	}

	err = newWordEntity.PointToCharacter(characterAggregate.GetWord())
	if err != nil {
		return nil, err
	}

	err = previousWordEntity.PointToWord(newWordEntity.GetWord())
	if err != nil {
		return nil, err
	}

	return newWordEntity, nil
}
