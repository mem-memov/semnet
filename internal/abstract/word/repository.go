package word

type Repository interface {
	Provide(word string) (Entity, error)
	Extract(entity Entity) (string, error)
	Fetch(phraseIdentifier uint) (Entity, error)
}
