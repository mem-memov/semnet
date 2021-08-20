package semnet

type words struct {
	storage    storage
	characters *characters
}

func newWords(storage storage, characters *characters) *words {
	return &words{
		storage:    storage,
		characters: characters,
	}
}

func (w *words) create(name string) (Word, error) {
	characters := make([]Character, len([]rune(name)))

	for i, r := range name {
		characters[i] = w.characters.create(r)
	}
}
