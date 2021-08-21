package bit

import "fmt"

type graph struct {
	storage storage
	isInitialized bool
}

func newGraph(storage storage) *graph {
	return &graph{
		storage: storage,
		isInitialized: false,
	}
}

func (g *graph) initialize() error {

	if g.isInitialized {
		return nil
	}

	hasZero, err := g.storage.Has(bitZeroNode)
	if err != nil {
		return err
	}

	if !hasZero {
		zeroNode, err := g.storage.Create()
		if err != nil {
			return err
		}

		if zeroNode != bitZeroNode {
			return fmt.Errorf("invalid zero node %d", zeroNode)
		}
	}

	hasOne, err := g.storage.Has(bitOneNode)
	if err != nil {
		return err
	}

	if !hasOne {
		oneNode, err := g.storage.Create()
		if err != nil {
			return err
		}

		if oneNode != bitOneNode {
			return fmt.Errorf("invalid one node %d", oneNode)
		}
	}

	g.isInitialized = true

	return nil
}