package remark

type chain struct {
	storage  storage
	entities *entities
}

func newChain(storage storage, entities *entities) *chain {
	return &chain{
		storage:  storage,
		entities: entities,
	}
}