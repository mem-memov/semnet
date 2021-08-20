package semnet

type Code struct {
	node  uint
	bits  bits
	codes codes
}

func newCode(node uint) Code {
	return Code{
		node: node,
	}
}

func (c Code) NextZero() (Code, error) {

}

func (c Code) NextOne() (Code, error) {

}

func (c Code) ToCharacter() (Character, error) {

}
