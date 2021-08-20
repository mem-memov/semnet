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

func (c *codes) createZero() (Code, error) {
	bit, err := c.bits.getZero()
	if err != nil {
		return Code{}, err
	}
	return newCode(bit), nil
}

func (c *codes) createOne() (Code, error) {

}
