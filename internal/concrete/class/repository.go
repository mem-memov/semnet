package class

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract/class"
	"github.com/mem-memov/semnet/internal/concrete/class/node"
)

const (
	bitIdentifier uint = iota + 1
	characterIdentifier
	wordIdentifier
	phraseIdentifier
	detailIdentifier
	remarkIdentifier
	factIdentifier
	storyIdentifier
)

type Repository struct {
	storage       storage
	isInitialized bool
	entity        Entity
}

var _ class.Repository = &Repository{}

func NewRepository(storage storage) *Repository {
	return &Repository{
		storage:       storage,
		isInitialized: false,
		entity:        Entity{},
	}
}

func (r *Repository) ProvideEntity() (class.Entity, error) {

	if !r.isInitialized {

		hasFirstNode, err := r.storage.Has(1)
		if err != nil {
			return Entity{}, err
		}

		if !hasFirstNode {
			result, err := r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != bitIdentifier {
				return Entity{}, fmt.Errorf("wrong bit class node %d", result)
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != characterIdentifier {
				return Entity{}, fmt.Errorf("wrong character class node %d", result)
			}

			err = r.storage.SetReference(bitIdentifier, characterIdentifier)
			if err != nil {
				return Entity{}, err
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != wordIdentifier {
				return Entity{}, fmt.Errorf("wrong word class node %d", result)
			}

			err = r.storage.SetReference(characterIdentifier, wordIdentifier)
			if err != nil {
				return Entity{}, err
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != phraseIdentifier {
				return Entity{}, fmt.Errorf("wrong phrase class node %d", result)
			}

			err = r.storage.SetReference(wordIdentifier, phraseIdentifier)
			if err != nil {
				return Entity{}, err
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != detailIdentifier {
				return Entity{}, fmt.Errorf("wrong detail class node %d", result)
			}

			err = r.storage.SetReference(phraseIdentifier, detailIdentifier)
			if err != nil {
				return Entity{}, err
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != remarkIdentifier {
				return Entity{}, fmt.Errorf("wrong position class node %d", result)
			}

			err = r.storage.SetReference(detailIdentifier, remarkIdentifier)
			if err != nil {
				return Entity{}, err
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != factIdentifier {
				return Entity{}, fmt.Errorf("wrong fact class node %d", result)
			}

			err = r.storage.SetReference(remarkIdentifier, factIdentifier)
			if err != nil {
				return Entity{}, err
			}

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != storyIdentifier {
				return Entity{}, fmt.Errorf("wrong story class node %d", result)
			}

			err = r.storage.SetReference(factIdentifier, storyIdentifier)
			if err != nil {
				return Entity{}, err
			}
		}

		r.entity = newEntity(
			node.NewBit(bitIdentifier, r.storage),
			node.NewCharacter(characterIdentifier, r.storage),
			node.NewWord(wordIdentifier, r.storage),
			node.NewPhrase(phraseIdentifier, r.storage),
			node.NewDetail(detailIdentifier, r.storage),
			node.NewRemark(remarkIdentifier, r.storage),
			node.NewFact(factIdentifier, r.storage),
			node.NewStory(storyIdentifier, r.storage),
		)
	}

	return r.entity, nil
}
