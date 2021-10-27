package fact

type Repository interface {
	CreateFirstUserStoryFact() (Entity, error)
}
