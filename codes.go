package semnet

type codes struct {
	bits *bits
}

func newCodes(bits *bits) *codes {
	return &codes{
		bits: bits,
	}
}
