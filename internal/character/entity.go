package character

import (
	"github.com/mem-memov/semnet/internal/character/node"
)

type Entity struct {
	codeNode      node.Code
	characterNode node.Character
	wordNode      node.Word
}

func newEntity(codeNode node.Code, characterNode node.Character, wordNode node.Word) Entity {
	return Entity{
		codeNode:      codeNode,
		characterNode: characterNode,
		wordNode:      wordNode,
	}
}
