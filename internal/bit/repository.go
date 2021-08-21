package bit

type Repository struct {
	storage  storage
	graph *graph
}

func NewRepository(storage storage) *Repository {
	return &Repository{
		storage:  storage,
		graph: newGraph(storage),
	}
}

func (r *Repository) create(value bool) (Entity, error) {

	err := r.graph.initialize()
	if err != nil {
		return Entity{}, err
	}

	if value {
		return newEntity(bitOneNode, r.storage), nil
	} else {
		return newEntity(bitZeroNode, r.storage), nil
	}
}
