package character

type Repository interface {
	Provide(integer rune) (Aggregate, error)
	Fetch(word uint) (Aggregate, error)
}
