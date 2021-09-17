package detail

import (
	"github.com/mem-memov/semnet/internal/detail/node"
	"github.com/mem-memov/semnet/internal/phrase"
)

type Entity struct {
	classNode  node.Class
	phraseNode node.Phrase
	remarkNode node.Remark
}

func newEntity(classNode node.Class, phraseNode node.Phrase, remarkNode node.Remark) Entity {
	return Entity{
		classNode:  classNode,
		phraseNode: phraseNode,
		remarkNode: remarkNode,
	}
}

func (e Entity) hasPhraseValue(phraseValue string) (bool, error) {

	hasPhraseValue, err := e.phraseNode.HasPhraseValue(phraseValue)
	if err != nil {
		return false, err
	}

	return hasPhraseValue, nil
}

func (e Entity) phraseValues() (string, string, error) {

	return e.phraseNode.PhraseValues()
}

func (e Entity) AddRemark(remarkIdentifier uint) error {

	return e.remarkNode.AddRemark(remarkIdentifier)
}

func (e Entity) GetObjectPhrase() (phrase.Entity, error) {

	return phrase.Entity{}, nil
}

func (e Entity) GetPropertyPhrase() (phrase.Entity, error) {

	return phrase.Entity{}, nil
}
