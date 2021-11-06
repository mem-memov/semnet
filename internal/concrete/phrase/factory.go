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

func (f *Factory) ProvideFirstEntity(classEntity abstractClass.Entity, wordEntity abstractWord.Aggregate) (abstractPhrase.Entity, error) {

	isBeginningOfPhrases, err := wordEntity.IsBeginningOfPhrases()
	if err != nil {
		return Entity{}, err
	}

	if !isBeginningOfPhrases {
		return f.phraseStorage.CreateEntity(classEntity, wordEntity)
	}

	word, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return nil, err
	}

	return f.phraseStorage.ReadEntityByWord(word)
}
