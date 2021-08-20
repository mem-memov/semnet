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
	for _, bitName := range fmt.Sprintf("%b", r) {
		switch bitName {
		case '0':

		case '1':

		default:
			return Character{}, fmt.Errorf("unexpected bit name: %c", bitName)
		}
	}
}
