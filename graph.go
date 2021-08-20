package semnet

import (
	"fmt"
	"strings"
)

type Graph struct {
	storage  storage
	zeroNode uint
	oneNode  uint
}

func NewGraph(storage storage) *Graph {
	return &Graph{
		storage:  storage,
		zeroNode: uint(1),
		oneNode:  uint(2),
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

func (g *Graph) GetZero() (Bit, error) {

	return newBit(g.zeroNode, g.storage), nil
}

func (g *Graph) GetOne() (Bit, error) {

	return newBit(g.oneNode, g.storage), nil
}

func (g *Graph) addAction(name string) (Action, error) {
	words := strings.Split(name, " ")
}
