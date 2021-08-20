package semnet

type codes struct {
	numbers *numbers
}

func newCodes(numbers *numbers) *codes {
	return &codes{
		numbers: numbers,
	}
}
