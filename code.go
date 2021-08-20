package semnet

type code struct {
	bit  bit
	bits  bits
	codes codes
}

func newCode(bit  bit) code {
	return code{
		bit: bit,
	}
}

func (c code) NextZero() (code, error) {

}

func (c code) NextOne() (code, error) {

}

func (c code) ToCharacter() (Character, error) {

}
