package detail

type Repository interface {
	Extend(objectIdentifier uint, property string) (Entity, error)
	Provide(object string, property string) (Entity, error)
	Fetch(remarkIdentifier uint) (Entity, error)
}
