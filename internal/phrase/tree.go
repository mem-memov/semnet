package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/word"
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

func (t *tree) provideRoot(wordEntity word.Entity) (Entity, error) {

	wordIdentifier, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	wordTargets, err := t.storage.ReadTargets(wordIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var phraseIdentifier uint
	var detailIdentifier uint

	switch len(wordTargets) {
	case 0:
		err = wordEntity.Mark(wordIdentifier)
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

		detailIdentifier, err = t.storage.Create()
		if err != nil {
			return Entity{}, err
		}

		err = t.storage.SetReference(phraseIdentifier, detailIdentifier)
		if err != nil {
			return Entity{}, err
		}
	case 1:
		if wordTargets[0] != wordEntity.PhraseIdentifier() {
			return Entity{}, fmt.Errorf("wrong target %d in detail tree at word %d", wordTargets[0], wordIdentifier)
		}

		_, wordIdentifier, err = t.storage.GetReference(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, phraseIdentifier, err = t.storage.GetReference(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in word tree at word %d", len(wordTargets), wordIdentifier)
	}

	return t.entities.create(wordIdentifier, phraseIdentifier, detailIdentifier), nil
}