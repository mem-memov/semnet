package phrase

import (
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Factory struct {
	storage abstractPhrase.Storage
}

var _ abstractPhrase.Factory = &Factory{}

func NewFactory(storage abstractPhrase.Storage) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) ProvideEntity(classEntity abstractClass.Entity, wordEntity abstractWord.Entity) (abstractPhrase.Entity, error) {

	hasWordSources, err := wordEntity.HasSingleTargetSources()
	if err != nil {
		return Entity{}, err
	}

	if !hasWordSources {
		return f.storage.CreateEntity(classEntity, wordEntity)
	}

	word, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return nil, err
	}

	return f.storage.ReadEntityByWord(word)
}
