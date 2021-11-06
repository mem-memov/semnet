package phrase

type Repository interface {
	Provide(words string) (Aggregate, error)
	Fetch(detail uint) (Aggregate, error)
}
