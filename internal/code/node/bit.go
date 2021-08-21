package node

type Bit uint

func NewBit(integer uint) Bit {
	return Bit(integer)
}

func (b Bit) ToInteger() uint {
	return uint(b)
}

func (b Bit) CreateCode(storage storage) (Code, error) {
	newInteger, err := storage.Create()
	if err != nil {
		return NewCode(0), err
	}

	codeNode := NewCode(newInteger)

	err = storage.SetReference(b.ToInteger(), codeNode.ToInteger())
	if err != nil {
		return NewCode(0), err
	}

	return codeNode, nil
}
