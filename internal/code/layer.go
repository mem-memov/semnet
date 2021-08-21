package code

import "github.com/mem-memov/semnet/internal/code/node"

type layer struct {
	storage storage
}

func newLayer(storage storage) *layer {
	return &layer{
		storage: storage,
	}
}

func (l *layer) createEntity(bitNode node.Bit) (Entity, error) {

	codeNode, err := bitNode.CreateCode(l.storage)
	if err != nil {
		return Entity{}, err
	}

	characterNode, err := codeNode.CreateCharacter(l.storage)
	if err != nil {
		return Entity{}, err
	}

	return newEntity(bitNode, codeNode, characterNode), nil
}
