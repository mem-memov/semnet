package story

import (
	"github.com/mem-memov/semnet/internal/concrete/detail"
	detail2 "github.com/mem-memov/semnet/internal/concrete/remark/detail"
	fact2 "github.com/mem-memov/semnet/internal/concrete/remark/fact"
	"github.com/mem-memov/semnet/internal/concrete/remark/remark"
	"github.com/mem-memov/semnet/internal/remark/node"
	"github.com/mem-memov/semnet/internal/remark/node/fact"
)

type entities struct {
	details *detail2.Factory
	remarks *remark.Remarks
	topics  *node.Topics
	facts   *fact2.Factory
}

func newEntities(storage storage, detailRepository *detail.Repository) *entities {
	return &entities{
		details: detail2.NewFactory(storage, detailRepository),
		remarks: remark.NewRemarks(storage),
		topics:  node.NewTopics(storage),
		facts:   fact.NewFacts(storage),
	}
}

func (e *entities) create(detailIdentifier uint, remarkIdentifier uint, topicIdentifier uint, factIdentifier uint) Entity {
	return newEntity(
		e.facts.Create(factIdentifier),
		e.topics.Create(topicIdentifier),
		e.remarks.Create(remarkIdentifier),
	)
}
