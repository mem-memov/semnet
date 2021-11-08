package phrase

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Factory struct {
	phraseStorage abstractPhrase.Storage
}

var _ abstractPhrase.Factory = &Factory{}

func NewFactory(phraseStorage abstractPhrase.Storage) *Factory {

	return &Factory{
		phraseStorage: phraseStorage,
	}
}

func (f *Factory) ProvideHeadEntity(classEntity abstractClass.Entity, wordAggregate abstractWord.Aggregate) (abstractPhrase.Entity, error) {

	hasTargetPhrase, err := wordAggregate.HasTargetPhrase()
	if err != nil {
		return nil, err
	}

	if !hasTargetPhrase {

		class, err := classEntity.CreatePhrase()
		if err != nil {
			return nil, err
		}

		phraseEntity, err := f.phraseStorage.CreateEntity(class)
		if err != nil {
			return nil, err
		}

		err = phraseEntity.PointToWord(wordAggregate.GetPhrase())
		if err != nil {
			return nil, err
		}

		err = wordAggregate.PointToPhrase(phraseEntity.GetWord())
		if err != nil {
			return nil, err
		}

		return phraseEntity, nil
	}

	word, err := wordAggregate.GetTargetPhrase()
	if err != nil {
		return nil, err
	}

	return f.phraseStorage.ReadEntityByWord(word)
}

func (f *Factory) CreateTailEntity(
	classEntity abstractClass.Entity,
	wordAggregate abstractWord.Aggregate,
	previousPhraseEntity abstractPhrase.Entity,
) (abstractPhrase.Entity, error) {

	class, err := classEntity.CreatePhrase()
	if err != nil {
		return nil, err
	}

	newPhraseEntity, err := f.phraseStorage.CreateEntity(class)
	if err != nil {
		return nil, err
	}

	err = newPhraseEntity.PointToWord(wordAggregate.GetPhrase())
	if err != nil {
		return nil, err
	}

	err = previousPhraseEntity.PointToPhrase(newPhraseEntity.GetPhrase())
	if err != nil {
		return nil, err
	}

	return newPhraseEntity, nil
}
