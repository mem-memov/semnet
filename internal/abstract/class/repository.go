package class

type Repository interface {
	ProvideEntity() (Entity, error)
}