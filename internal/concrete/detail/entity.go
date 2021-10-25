package detail

import (
	node2 "github.com/mem-memov/semnet/internal/concrete/detail/node"
)

type Entity struct {
	classNode  node2.Class
	phraseNode node2.Phrase
	remarkNode node2.Remark
}

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

func (e Entity) AddRemark(remarkIdentifier uint) error {

	return e.remarkNode.AddRemark(remarkIdentifier)
}

func (e Entity) GetObjectAndProperty() (string, string, error) {
	return e.phraseNode.PhraseValues()
}
