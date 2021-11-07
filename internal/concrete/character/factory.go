package character

import (
	abstractBit "github.com/mem-memov/semnet/internal/abstract/bit"
	abstractCharacter "github.com/mem-memov/semnet/internal/abstract/character"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
)

type Factory struct {
	characterStorage abstractCharacter.Storage
}

var _ abstractCharacter.Factory = &Factory{}

func NewFactory(
	characterStorage abstractCharacter.Storage,
) *Factory {

	return &Factory{
		characterStorage: characterStorage,
	}
}

func (f *Factory) ProvideFirstEntity(
	classEntity abstractClass.Entity,
	bitEntity abstractBit.Entity,
) (abstractCharacter.Entity, error) {

	hasCharacterSources, err := bitEntity.IsBeginningOfCharacters()
	if err != nil {
		return Entity{}, err
	}

	if !hasCharacterSources {
		return f.characterStorage.CreateEntity(classEntity, bitEntity)
	}

	bit, err := bitEntity.ProvideSingleTarget()
	if err != nil {
		return nil, err
	}

	return f.characterStorage.ReadEntityByBit(bit)
}
