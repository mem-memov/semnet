package word

type Repository interface {
	Provide(word string) (interface{}, error)         // use Entity after refactoring this package
	Extract(entity interface{}) (string, error)       // use Entity after refactoring this package
	Fetch(phraseIdentifier uint) (interface{}, error) // use Entity after refactoring this package
}
