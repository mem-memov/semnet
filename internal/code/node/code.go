package node

type Code uint

func NewCode(integer uint) Code {
	return Code(integer)
}

func (c Code) ToInteger() uint {
	return uint(c)
}

func (c Code) IsEmpty() bool {
	return c == 0
}

func (c Code) CreateCharacter(storage storage) (Character, error) {

	newInteger, err := storage.Create()
	if err != nil {
		return NewCharacter(0), err
	}

	characterNode := NewCharacter(newInteger)

	err = storage.SetReference(c.ToInteger(), characterNode.ToInteger())
	if err != nil {
		return NewCharacter(0), err
	}

	return characterNode, nil
}

func (c Code) ReadTargets()
