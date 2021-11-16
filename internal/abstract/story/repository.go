package story

type Repository interface {
	CreateStory() (Entity, error)
	FetchByFact(factIdentifier uint) (Entity, error)
	FetchByPosition(positionIdentifier uint) (Entity, error)
	FetchByTree(treeIdentifier uint) (Entity, error)
}
