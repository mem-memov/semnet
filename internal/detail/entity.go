package detail

import (
	"github.com/mem-memov/semnet/internal/detail/node"
)

type Entity struct {
	phraseNode node.Phrase
	remarkNode node.Remark
}

func newEntity(phraseNode node.Phrase, remarkNode node.Remark) Entity {
	return Entity{
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
