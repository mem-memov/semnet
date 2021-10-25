package remark

import (
	abstractRemarkClass "github.com/mem-memov/semnet/internal/abstract/remark/class"
	abstractRemarkDetail "github.com/mem-memov/semnet/internal/abstract/remark/detail"
	abstractRemarkFact "github.com/mem-memov/semnet/internal/abstract/remark/fact"
	abstractRemarkRemark "github.com/mem-memov/semnet/internal/abstract/remark/remark"
)

type Entity struct {
	classNode  abstractRemarkClass.Node
	detailNode abstractRemarkDetail.Node
	remarkNode abstractRemarkRemark.Node
	factNode   abstractRemarkFact.Node
}

func newEntity(
	classNode abstractRemarkClass.Node,
	detailNode abstractRemarkDetail.Node,
	remarkNode abstractRemarkRemark.Node,
	factNode abstractRemarkFact.Node,
) Entity {
	return Entity{
		classNode: classNode,
		detailNode: detailNode,
		remarkNode: remarkNode,
		factNode:   factNode,
	}
}