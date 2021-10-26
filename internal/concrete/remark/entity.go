package remark

import (
	abstractRemarkClass "github.com/mem-memov/semnet/internal/abstract/remark/class"
	abstractRemarkDetail "github.com/mem-memov/semnet/internal/abstract/remark/detail"
	abstractRemarkFact "github.com/mem-memov/semnet/internal/abstract/remark/fact"
	abstractRemarkPosition "github.com/mem-memov/semnet/internal/abstract/remark/position"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
)

type Entity struct {
	classNode  abstractRemarkClass.Node
	detailNode   abstractRemarkDetail.Node
	positionNode abstractRemarkPosition.Node
	factNode     abstractRemarkFact.Node
}

var _ abstractRemark.Entity = Entity{}

func newEntity(
	classNode abstractRemarkClass.Node,
	detailNode abstractRemarkDetail.Node,
	positionNode abstractRemarkPosition.Node,
	factNode abstractRemarkFact.Node,
) Entity {
	return Entity{
		classNode:    classNode,
		detailNode:   detailNode,
		positionNode: positionNode,
		factNode:     factNode,
	}
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.factNode.Mark(sourceIdentifier)
}