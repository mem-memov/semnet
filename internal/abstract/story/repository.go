package story

type Repository interface {
	CreateFirstUserStory() (Entity, error)
	FetchByFact(factIdentifier uint) (Entity, error)
}
