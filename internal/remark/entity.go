package remark

import (
	"github.com/mem-memov/semnet/internal/detail"
	"github.com/mem-memov/semnet/internal/remark/node"
)

type Entity struct {
	detailNode node.Detail
	remarkNode node.Remark
	topicNode  node.Topic
	factNode   node.Fact
}

func newEntity(detailNode node.Detail, remarkNode node.Remark, topicNode node.Topic, factNode node.Fact) Entity {
	return Entity{
		detailNode: detailNode,
		remarkNode: remarkNode,
		topicNode:  topicNode,
		factNode:   factNode,
	}
}

func (e Entity) GetDetail() (detail.Entity, error) {

	return detail.Entity{}, nil
}
