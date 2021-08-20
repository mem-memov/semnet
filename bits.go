package semnet

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

func (b *bits) GetZero() (Bit, error) {

	return newBit(b.zeroNode, b.storage), nil
}

func (b *bits) GetOne() (Bit, error) {

	return newBit(b.oneNode, b.storage), nil
}
