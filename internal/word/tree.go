package word

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/character"
)

type tree struct {
	storage  storage
	entities *entities
}

func newTree(storage storage, entities *entities) *tree {
	return &tree{
		storage:  storage,
		entities: entities,
	}
}

func (t *tree) provideRoot(characterEntity character.Entity) (Entity, error) {

	characterIdentifier, err := characterEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	characterTargets, err := t.storage.ReadTargets(characterIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var classIdentifier uint
	var wordIdentifier uint
	var phraseIdentifier uint

	switch len(characterTargets) {
	case 0:
		err = characterEntity.Mark(characterIdentifier)
		if err != nil {
			return Entity{}, err
		}

		wordIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(characterIdentifier, wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

		phraseIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(wordIdentifier, phraseIdentifier)
		if err != nil {
			return Entity{}, err
		}

		entity, err := t.entities.createAndAddClass(characterIdentifier, wordIdentifier, phraseIdentifier)
		if err != nil {
			return Entity{}, err
		}

		return entity, nil
	case 1:
		if characterTargets[0] != characterEntity.WordIdentifier() {
			return Entity{}, fmt.Errorf("wrong target %d in word tree at character %d", characterTargets[0], characterIdentifier)
		}

		classIdentifier, wordIdentifier, err = t.storage.GetReference(characterIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, phraseIdentifier, err = t.storage.GetReference(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

		return t.entities.create(classIdentifier, characterIdentifier, wordIdentifier, phraseIdentifier), nil
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in word tree at character %d", len(characterTargets), characterIdentifier)
	}
}
