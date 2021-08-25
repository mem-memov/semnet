package phrase

import (
	"github.com/mem-memov/semnet/internal/phrase/node"
)

type Entity struct {
	wordNode      node.Word
	phraseNode    node.Phrase
	detailNode node.Detail
}

func newEntity(wordNode node.Word, phraseNode node.Phrase, detailNode node.Detail) Entity {
	return Entity{
		wordNode:      wordNode,
		phraseNode:    phraseNode,
		detailNode: detailNode,
	}
}