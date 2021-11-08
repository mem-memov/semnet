package phrase

type Storage interface {
	CreateEntity(class uint) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByWord(word uint) (Entity, error)
	ReadEntityByPhrase(phrase uint) (Entity, error)
	ReadEntityByDetail(detail uint) (Entity, error)
}
