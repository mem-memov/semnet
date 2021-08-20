package semnet

type codes struct {
	storage storage
	bits *bits
}

func newCodes(storage storage, bits *bits) *codes {
	return &codes{
		storage: storage,
		bits: bits,
	}
}
