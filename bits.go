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
