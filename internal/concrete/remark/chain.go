package remark

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type chain struct {
	storage abstract.Storage
	factory abstractRemark.Factory
}

func newChain(storage abstract.Storage, factory abstractRemark.Factory) *chain {
	return &chain{
		storage: storage,
		factory: factory,
	}
}

func (c *chain) CreateFirstLink(detailEntity abstractDetail.Entity) (abstractRemark.Entity, error) {

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

	return c.factory.CreateNewEntity(detailIdentifier, remarkIdentifier, factIdentifier), nil
}
