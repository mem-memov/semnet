package word

type Storage interface {
	CreateEntity(class uint) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByCharacter(character uint) (Entity, error)
	ReadEntityByWord(word uint) (Entity, error)
	ReadEntityByPhrase(phrase uint) (Entity, error)
}
