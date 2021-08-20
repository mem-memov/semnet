package semnet

type Code struct {
	bit  bit
	bits  bits
	codes codes
}

func newCode(bit  bit) Code {
	return Code{
		bit: bit,
	}
}

func (c Code) NextZero() (Code, error) {

}

func (c Code) NextOne() (Code, error) {

}

func (c Code) ToCharacter() (Character, error) {

}
