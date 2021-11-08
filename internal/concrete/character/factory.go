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

func (f *Factory) ProvideHeadEntity(
	classEntity abstractClass.Entity,
	bitEntity abstractBit.Entity,
) (abstractCharacter.Entity, error) {

	hasSingleTargetCharacter, err := bitEntity.HasTargetCharacter()
	if err != nil {
		return nil, err
	}

	if !hasSingleTargetCharacter {
		class, err := classEntity.CreateCharacter()
		if err != nil {
			return nil, err
		}

		characterEntity, err := f.characterStorage.CreateEntity(class)
		if err != nil {
			return nil, err
		}

		err = characterEntity.PointToBit(bitEntity.GetCharacter())
		if err != nil {
			return nil, err
		}

		err = bitEntity.PointToCharacter(characterEntity.GetBit())
		if err != nil {
			return nil, err
		}

		return characterEntity, nil
	}

	bit, err := bitEntity.GetTargetCharacter()
	if err != nil {
		return nil, err
	}

	return f.characterStorage.ReadEntityByBit(bit)
}

func (f *Factory) CreateTailEntity(
	classEntity abstractClass.Entity,
	bitEntity abstractBit.Entity,
	previousCharacterEntity abstractCharacter.Entity,
) (abstractCharacter.Entity, error) {

	class, err := classEntity.CreateCharacter()
	if err != nil {
		return nil, err
	}

	newCharacterEntity, err := f.characterStorage.CreateEntity(class)
	if err != nil {
		return nil, err
	}

	err = newCharacterEntity.PointToBit(bitEntity.GetCharacter())
	if err != nil {
		return nil, err
	}

	err = previousCharacterEntity.PointToCharacter(newCharacterEntity.GetCharacter())
	if err != nil {
		return nil, err
	}

	return newCharacterEntity, nil
}
