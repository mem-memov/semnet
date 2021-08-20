package semnet

type bit struct {
	node    uint
	storage storage
}

func newBit(node uint, storage storage) bit {
	return bit{
		node:    node,
		storage: storage,
	}
}
