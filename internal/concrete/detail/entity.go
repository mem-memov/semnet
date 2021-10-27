package detail

import (
	abstractDetail "github.com/mem-memov/semnet/internal/abstract/detail"
	abstractRemark "github.com/mem-memov/semnet/internal/abstract/remark"
	node2 "github.com/mem-memov/semnet/internal/concrete/detail/node"
)

type Entity struct {
	classNode  node2.Class
	phraseNode node2.Phrase
	remarkNode node2.Remark
}

var _ abstractDetail.Entity = Entity{}

func newEntity(classNode node2.Class, phraseNode node2.Phrase, remarkNode node2.Remark) Entity {
	return Entity{
		classNode:  classNode,
		phraseNode: phraseNode,
		remarkNode: remarkNode,
	}
}

func (e Entity) HasPhraseValue(phraseValue string) (bool, error) {

	hasPhraseValue, err := e.phraseNode.HasPhraseValue(phraseValue)
	if err != nil {
		return false, err
	}

	return hasPhraseValue, nil
}

func (e Entity) PointToRemark(remark abstractRemark.Entity) error {

	return e.remarkNode.PointToRemark(remark.GetDetail())
}

func (e Entity) GetObjectAndProperty() (string, string, error) {
	return e.phraseNode.PhraseValues()
}
