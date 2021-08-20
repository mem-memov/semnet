package semnet

type Character struct {
	node uint
}

func newCharacter(node uint) Character {
	return Character{
		node: node,
	}
}

func (c Character) NextCode(code rune) (Code, error) {

}

func (c Character) ToWord() (Word, error) {

}
