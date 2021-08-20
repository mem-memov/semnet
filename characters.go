package semnet

import "fmt"

type characters struct {
	storage storage
	codes *codes
}

func newCharacters(storage storage, codes *codes) *characters {
	return &characters{
		storage: storage,
		codes: codes,
	}
}

func (c *characters) create(r rune) (Character, error) {

	var code code
	var err error

	for i, bitName := range fmt.Sprintf("%b", r) {
		switch bitName {
		case '0':
			if i == 0 {
				code, err = c.codes.createZero()
				if err != nil {
					return Character{}, err
				}
			} else {
				code, err = code.NextZero()
				if err != nil {
					return Character{}, err
				}
			}
		case '1':
			if i == 0 {
				code, err = c.codes.createOne()
				if err != nil {
					return Character{}, err
				}
			} else {
				code, err = code.NextOne()
				if err != nil {
					return Character{}, err
				}
			}
		default:
			return Character{}, fmt.Errorf("unexpected bit name: %c", bitName)
		}
	}

	return newCharacter(code), nil
}
