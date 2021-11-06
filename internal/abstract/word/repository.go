package word

type Repository interface {
	Provide(word string) (Aggregate, error)
	Fetch(phraseIdentifier uint) (Aggregate, error)
}
