package story

type Repository interface {
	CreateFirstUserStory() (Entity, error)
}
