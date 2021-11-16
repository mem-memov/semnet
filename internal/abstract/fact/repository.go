package fact

type Repository interface {
	CreateFirstStoryFact() (Aggregate, error)
	FetchByRemark(remarkIdentifier uint) (Aggregate, error)
	CreateNextFact(remarkIdentifier uint) (Aggregate, error)
	CreateChildStoryFact(remarkIdentifier uint) (Aggregate, error)
}
