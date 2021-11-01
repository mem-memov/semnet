package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type tree struct {
	storage         abstract.Storage
	classRepository abstractClass.Repository
}

var _ abstractPhrase.Tree = &tree{}

func newTree(storage abstract.Storage, classRepository abstractClass.Repository) *tree {
	return &tree{
		storage:         storage,
		classRepository: classRepository,
	}
}

func (t *tree) ProvideRoot(wordEntity abstractWord.Entity) (abstractPhrase.Entity, error) {

	wordIdentifier, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	wordTargets, err := t.storage.ReadTargets(wordIdentifier)
	if err != nil {
		return Entity{}, err
	}

	switch len(wordTargets) {
	case 0:
		err = wordEntity.Mark(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

		class, err := t.classRepository.ProvideEntity()
		if err != nil {
			return Entity{}, err
		}

		return createEntity(t.storage, class)

	case 1:
		if wordTargets[0] != wordEntity.PhraseIdentifier() {
			return Entity{}, fmt.Errorf("wrong target %d in detail tree at word %d", wordTargets[0], wordIdentifier)
		}

		return readEntityByWord(t.storage, wordIdentifier)
	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in word tree at word %d", len(wordTargets), wordIdentifier)
	}
}
