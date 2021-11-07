package character

type Aggregate interface {
	Extract() (rune, error)
}
