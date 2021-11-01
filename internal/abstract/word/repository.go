package word

type Repository interface {
	Provide(word string) (Entity, error)
}
