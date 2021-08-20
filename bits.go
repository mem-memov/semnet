package semnet

import "fmt"

type bits struct {
	storage  storage
	zeroNode uint
	oneNode  uint
}

func newBits(storage storage) *bits {
	return &bits{
		storage:  storage,
		zeroNode: uint(1),
		oneNode:  uint(2),
	}
}

func (b *bits) initialize() error {

	hasZero, err := b.storage.Has(b.zeroNode)
	if err != nil {
		return err
	}

	if !hasZero {
		zeroNode, err := b.storage.Create()
		if err != nil {
			return err
		}

		if zeroNode != b.zeroNode {
			return fmt.Errorf("invalid zero node %d", zeroNode)
		}
	}

	hasOne, err := b.storage.Has(b.oneNode)
	if err != nil {
		return err
	}

	if !hasOne {
		oneNode, err := b.storage.Create()
		if err != nil {
			return err
		}

		if oneNode != b.oneNode {
			return fmt.Errorf("invalid one node %d", oneNode)
		}
	}

	return nil
}

func (b *bits) getZero() (bit, error) {

	return newBit(b.zeroNode, b.storage), nil
}

func (b *bits) getOne() (bit, error) {

	return newBit(b.oneNode, b.storage), nil
}
