package character

import (
	code2 "github.com/mem-memov/semnet/internal/code"
	"github.com/mem-memov/semnet/internal/word"
)

type Character struct {
	code code2.code
}

func newCharacter(code code2.code) Character {
	return Character{
		code: code,
	}
}

func (c Character) NextCode(code rune) (code2.code, error) {

}

func (c Character) ToWord() (word.Word, error) {

}
