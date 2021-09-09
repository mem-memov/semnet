package story

import (
	"github.com/mem-memov/semnet/internal/detail"
	"github.com/mem-memov/semnet/internal/remark/node"
)

type entities struct {
	details *node.Details
	remarks *node.Remarks
	topics  *node.Topics
	facts   *node.Facts
}

func newEntities(storage storage, detailRepository *detail.Repository) *entities {
	return &entities{
		details: node.NewDetails(storage, detailRepository),
		remarks: node.NewRemarks(storage),
		topics:  node.NewTopics(storage),
		facts:   node.NewFacts(storage),
	}
}

func (e *entities) create(detailIdentifier uint, remarkIdentifier uint, topicIdentifier uint, factIdentifier uint) Entity {
	return newEntity(
		e.details.Create(detailIdentifier),
		e.remarks.Create(remarkIdentifier),
		e.topics.Create(topicIdentifier),
		e.facts.Create(factIdentifier),
	)
}
