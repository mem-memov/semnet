package fact

type Repository interface {
	CreateFirstUserStoryFact() (Aggregate, error)
	FetchByRemark(remarkIdentifier uint) (Aggregate, error)
}
