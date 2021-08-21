package word

import (
	"github.com/mem-memov/semnet/internal/character"
)

type Word struct {
}

func newWord() Word {
	return Word{}
}

func (w Word) NextCharacter() (character.Character, error) {

}
