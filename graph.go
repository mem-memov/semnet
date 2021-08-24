package semnet

import (
	"github.com/mem-memov/semnet/internal"
	"github.com/mem-memov/semnet/internal/action"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/word"
)

type Graph struct {
	bits    *bit.bits
	actions *action.actions
}

func NewGraph(storage internal.storage) *Graph {
	bits := bit.newBits(storage)
	codes := code.newCodes(storage, bits)
	characters := word.newCharacters(storage, codes)
	words := word.newWords(storage, characters)
	actions := action.newActions(storage, words)

	return &Graph{
		bits:    bits,
		actions: actions,
	}
}

func (g *Graph) InitializeBits() error {
	return g.bits.initialize()
}

func (g *Graph) AddAction(name string) (action.Action, error) {
	return g.actions.create(name)
}
