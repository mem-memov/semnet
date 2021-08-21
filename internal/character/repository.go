package character

import (
	"fmt"
	"github.com/mem-memov/semnet"
	"github.com/mem-memov/semnet/internal/code"
)

type Repository struct {
	storage semnet.storage
	codes   *code.codes
}

func NewRepository(storage storage, codeRepository *code.Repository) *Repository {
	return &Repository{
		storage: storage,
		codes:   codeRepository,
	}
}

func (c *Repository) create(rune rune) (Entity, error) {

	code, err := c.codes.create(int32(rune))
	if err != nil {
		return Entity{}, err
	}

	var code code
	var err error

	for i, bitName := range fmt.Sprintf("%b", r) {

		switch bitName {
		case '0':
			if i == 0 {
				code, err = c.codes.createZero()
				if err != nil {
					return Entity{}, err
				}
			} else {
				code, err = code.NextZero()
				if err != nil {
					return Entity{}, err
				}
			}
		case '1':
			if i == 0 {
				code, err = c.codes.createOne()
				if err != nil {
					return Entity{}, err
				}
			} else {
				code, err = code.NextOne()
				if err != nil {
					return Entity{}, err
				}
			}
		default:
			return Entity{}, fmt.Errorf("unexpected bit name: %c", bitName)
		}
	}

	return newCharacter(code), nil
}
