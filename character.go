package semnet

type Character struct {
	code code
}

func newCharacter(code code) Character {
	return Character{
		code: code,
	}
}

func (c Character) NextCode(code rune) (code, error) {

}

func (c Character) ToWord() (Word, error) {

}
