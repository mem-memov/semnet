package phrase

import (
	"fmt"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	"github.com/mem-memov/semnet/internal/concrete/word"
)

type tree struct {
	storage  storage
	entities abstractPhrase.Entities
}

func newTree(storage storage, entities abstractPhrase.Entities) abstractPhrase.Tree {
	return &tree{
		storage:  storage,
		entities: entities,
	}
}

func (t *tree) ProvideRoot(wordEntity word.Entity) (abstractPhrase.Entity, error) {

	wordIdentifier, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	wordTargets, err := t.storage.ReadTargets(wordIdentifier)
	if err != nil {
		return Entity{}, err
	}

	var classIdentifier uint
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

		entity, err := t.entities.CreateAndAddClass(wordIdentifier, phraseIdentifier, detailIdentifier)
		if err != nil {
			return Entity{}, err
		}

		return entity, nil
	case 1:
		if wordTargets[0] != wordEntity.PhraseIdentifier() {
			return Entity{}, fmt.Errorf("wrong target %d in detail tree at word %d", wordTargets[0], wordIdentifier)
		}

		classIdentifier, wordIdentifier, err = t.storage.GetReference(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

		_, phraseIdentifier, err = t.storage.GetReference(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

		return t.entities.Create(classIdentifier, wordIdentifier, phraseIdentifier, detailIdentifier), nil
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in word tree at word %d", len(wordTargets), wordIdentifier)
	}
}
