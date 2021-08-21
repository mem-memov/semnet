package character

import (
	code2 "github.com/mem-memov/semnet/internal/code"
	"github.com/mem-memov/semnet/internal/word"
)

type Entity struct {
	code code2.code
}

func newCharacter(code code2.code) Entity {
	return Entity{
		code: code,
	}
}

func (c Entity) NextCode(code rune) (code2.code, error) {

}

func (c Entity) ToWord() (word.Word, error) {

}
