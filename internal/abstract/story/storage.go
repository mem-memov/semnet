package story

type Storage interface {
	CreateEntity(class uint) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByFact(fact uint) (Entity, error)
	ReadEntityByPosition(position uint) (Entity, error)
	ReadEntityByUser(user uint) (Entity, error)
}


