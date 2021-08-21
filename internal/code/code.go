package code

import (
	"github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/character"
)

type code struct {
	bit           semnet.bit
	codeNode      uint
	characterNode uint
	bits          semnet.bits
	codes         codes
	storage       semnet.storage
}

func newCode(bit semnet.bit, codeNode uint, characterNode uint) code {
	return code{
		bit:           bit,
		codeNode:      codeNode,
		characterNode: characterNode,
	}
}

func (c code) NextZero() (code, error) {

}

func (c code) NextOne() (code, error) {

}

func (c code) ToCharacter() (character.Character, error) {

}
