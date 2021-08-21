package node

type Character uint

func NewCharacter(integer uint) Character {
	return Character(integer)
}

func (c Character) ToInteger() uint {
	return uint(c)
}

func (c Character) IsEmpty() bool {
	return c == 0
}
