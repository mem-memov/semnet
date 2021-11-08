package character

type Storage interface {
	CreateEntity(class uint) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByBit(bit uint) (Entity, error)
	ReadEntityByCharacter(character uint) (Entity, error)
	ReadEntityByWord(word uint) (Entity, error)
}
