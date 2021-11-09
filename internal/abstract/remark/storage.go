package remark

type Storage interface {
	CreateEntity(class uint) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByDetail(detail uint) (Entity, error)
	ReadEntityByPosition(position uint) (Entity, error)
	ReadEntityByFact(fact uint) (Entity, error)
}
