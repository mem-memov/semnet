package semnet

type words struct {
	characters *characters
}

func newWords(characters *characters) *words {
	return &words{
		characters: characters,
	}
}

func (w *words) create(name string) (Word, error) {
	characters := make([]Character, len([]rune(name)))

	for i, r := range name {
		characters[i] = w.characters.create(r)
	}
}
