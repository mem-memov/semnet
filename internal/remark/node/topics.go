package node

type Topics struct {
	storage storage
}

func NewTopics(storage storage) *Topics {
	return &Topics{
		storage: storage,
	}
}

func (t *Topics) Create(identifier uint) Topic {
	return newTopic(identifier, t.storage)
}