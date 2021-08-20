package semnet

type Bit struct {
	node  uint
	bits  bits
	codes codes
}

func newBit(node uint, bits bits, codes codes) Bit {
	return Bit{
		node:  node,
		bits:  bits,
		codes: codes,
	}
}

func ToCode() (Code, error) {

	return newCode()
}
