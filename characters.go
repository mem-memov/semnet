package semnet

import "fmt"

type characters struct {
	codes *codes
}

func newCharacters(codes *codes) *characters {
	return &characters{
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
