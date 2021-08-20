package semnet

import (
	"fmt"
)

type Graph struct {
	storage  storage
	zeroNode uint
	oneNode  uint
	actions *actions
}

func NewGraph(storage storage) *Graph {
	bits := newBits(storage)
	codes := newCodes(storage, bits)
	characters := newCharacters(storage, codes)
	words := newWords(storage, characters)
	actions := newActions(storage, words)

	return &Graph{
		storage:  storage,
		zeroNode: uint(1),
		oneNode:  uint(2),
		actions: actions,
	}
}

func (g *Graph) InitializeBits() error {

	hasZero, err := g.storage.Has(g.zeroNode)
	if err != nil {
		return err
	}

	if !hasZero {
		zeroNode, err := g.storage.Create()
		if err != nil {
			return err
		}

		if zeroNode != g.zeroNode {
			return fmt.Errorf("invalid zero node %d", zeroNode)
		}
	}

	hasOne, err := g.storage.Has(g.oneNode)
	if err != nil {
		return err
	}

	if !hasOne {
		oneNode, err := g.storage.Create()
		if err != nil {
			return err
		}

		if oneNode != g.oneNode {
			return fmt.Errorf("invalid one node %d", oneNode)
		}
	}

	return nil
}

func (g *Graph) addAction(name string) (Action, error) {
	return g.actions.create(name)
}
