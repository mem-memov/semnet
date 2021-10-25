package fact

import (
	node2 "github.com/mem-memov/semnet/internal/concrete/fact/node"
	"github.com/mem-memov/semnet/internal/remark"
)

type Entity struct {
	classNode  node2.Class
	remarkNode node2.Remark
	factNode   node2.Fact
	storyNode  node2.Story
}

func (e Entity) IdentifierForRemark() uint {
	return e.remarkNode.Identifier()
}

func (e Entity) AddRemark(object string, property string) (remark.Entity, error) {

	return remark.Entity{}, nil
}

func (e Entity) GetFact() (Fact, error) {

	return nil, nil
}
