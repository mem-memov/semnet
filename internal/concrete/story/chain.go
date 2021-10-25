package story

import (
	"github.com/mem-memov/semnet/internal/concrete/detail"
)

type chain struct {
	storage  storage
	entities *entities
}

func newChain(storage storage, entities *entities) *chain {
	return &chain{
		storage:  storage,
		entities: entities,
	}
}

func (c *chain) createFirstLink(detailEntity detail.Entity) (Entity, error) {

	detailIdentifier, err := c.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = detailEntity.AddRemark(detailIdentifier)
	if err != nil {
		return Entity{}, err
	}

	remarkIdentifier, err := c.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = c.storage.SetReference(detailIdentifier, remarkIdentifier)
	if err != nil {
		return Entity{}, err
	}

	topicIdentifier, err := c.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = c.storage.SetReference(remarkIdentifier, topicIdentifier)
	if err != nil {
		return Entity{}, err
	}

	factIdentifier, err := c.storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = c.storage.SetReference(remarkIdentifier, factIdentifier)
	if err != nil {
		return Entity{}, err
	}

	return c.entities.create(detailIdentifier, remarkIdentifier, topicIdentifier, factIdentifier), nil
}
