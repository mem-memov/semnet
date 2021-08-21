package character

import (
	"github.com/mem-memov/semnet/internal/character/node"
	"github.com/mem-memov/semnet/internal/code"
)

type entities struct {
	codes      *node.Codes
	characters *node.Characters
	words      *node.Words
}

func newEntities(storage storage, codeRepository *code.Repository) *entities {
	return &entities{
		codes:      node.NewCodes(storage, codeRepository),
		characters: node.NewCharacters(storage),
		words:      node.NewWords(storage),
	}
}

func (e *entities) create(codeIdentifier uint, characterIdentifier uint, wordIdentifier uint) Entity {
	return newEntity(
		e.codes.Create(codeIdentifier),
		e.characters.Create(characterIdentifier),
		e.words.Create(wordIdentifier),
	)
}
