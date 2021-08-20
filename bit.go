package semnet

type bit struct {
	node  uint
	bits  bits
	codes codes
}

func newBit(node uint, bits bits, codes codes) bit {
	return bit{
		node:  node,
		bits:  bits,
		codes: codes,
	}
}

func ToCode() (Code, error) {

	return newCode(), nil
}
