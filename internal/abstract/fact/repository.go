package fact

type Repository interface {
	CreateFirstUserStoryFact() (Entity, error)
	FetchByRemark(remarkIdentifier uint) (Entity, error)
}
