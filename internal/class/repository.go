package class

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/class/node"
)

type Repository struct {
	storage       storage
	isInitialized bool
	entity        Entity
}

const (
	bitIdentifier uint = iota + 1
	characterIdentifier
	wordIdentifier
	phraseIdentifier
	detailIdentifier
	remarkIdentifier
	factIdentifier
	storyIdentifier
	userIdentifier
)

func NewRepository(storage storage) *Repository {
	return &Repository{
		storage:       storage,
		isInitialized: false,
		entity:        Entity{},
	}
}

func (r *Repository) ProvideEntity() (Entity, error) {

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
				return Entity{}, fmt.Errorf("wrong remark class node %d", result)
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

			result, err = r.storage.Create()
			if err != nil {
				return Entity{}, err
			}

			if result != userIdentifier {
				return Entity{}, fmt.Errorf("wrong user class node %d", result)
			}

			err = r.storage.SetReference(storyIdentifier, userIdentifier)
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
			node.NewUser(userIdentifier, r.storage),
		)
	}

	return r.entity, nil
}
