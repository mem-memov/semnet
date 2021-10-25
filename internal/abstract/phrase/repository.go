package phrase

type Repository interface {
	Provide(words string) (Entity, error)
	Extract(entity Entity) (string, error)
	Fetch(detailIdentifier uint) (Entity, error)
}