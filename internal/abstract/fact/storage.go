package fact

type Storage interface {
	CreateEntity(class uint) (Entity, error)
	ReadEntityByClass(class uint) (Entity, error)
	ReadEntityByRemark(remark uint) (Entity, error)
	ReadEntityByPosition(position uint) (Entity, error)
	ReadEntityByStory(story uint) (Entity, error)
}
