package character

type Repository interface {
	Provide(integer rune) (interface{}, error)
	Extract(entity interface{}) (rune, error)
	Fetch(wordIdentifier uint) (interface{}, error)
}
