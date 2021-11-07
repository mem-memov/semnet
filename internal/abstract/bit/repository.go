package bit

type Repository interface {
	Provide(value bool) (Entity, error)
	Fetch(identifier uint) (Entity, error)
}
