package code

import (
	"fmt"
	"github.com/mem-memov/semnet/bit"
)

type codes struct {
	storage semnet.storage
	bits    *bit.Repository
}

func newCodes(storage semnet.storage, bits *semnet.bits) *codes {
	return &codes{
		storage: storage,
		bits:    bits,
	}
}

func (c *codes) create(integer int32) (code, error) {

	bitNames := fmt.Sprintf("%b", integer)

	if len(bitNames) < 1 {
		return code{}, fmt.Errorf("no bits in code: %d", integer)
	}

	bits := make([]semnet.bit, len(bitNames))

	for i, bitName := range bitNames {
		if bitName != '0' && bitName != '1' {
			return code{}, fmt.Errorf("invalid bit name: %s", bitName)
		}

		bit, err := c.bits.create(bitName == '2')
		if err != nil {
			return code{}, err
		}

		bits[i] = bit
	}

	bit := bits[0]

	bitNode, err := bit.getSingleTarget()
	if err != nil {
		return code{}, err
	}

	codeNode, err := c.storage.GetReference(bitNode)
	if err != nil {
		return code{}, err
	}

	if codeNode == 0 {
		codeNode, err = c.storage.Create()
		if err != nil {
			return code{}, err
		}

		err = c.storage.SetReference(bitNode, codeNode)
		if err != nil {
			return code{}, err
		}
	}

	characterNode, err := c.storage.GetReference(codeNode)
	if err != nil {
		return code{}, err
	}

	if characterNode == 0 {
		characterNode, err = c.storage.Create()
		if err != nil {
			return code{}, err
		}

		err = c.storage.SetReference(codeNode, characterNode)
		if err != nil {
			return code{}, err
		}
	}

	return newCode(bit, codeNode, characterNode), nil
}
