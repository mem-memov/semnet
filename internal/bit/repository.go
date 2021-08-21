package bit

type Repository struct {
	storage storage
	layer   *layer
}

func NewRepository(storage storage) *Repository {
	return &Repository{
		storage: storage,
		layer:   newLayer(storage),
	}
}

func (r *Repository) Create(value bool) (Entity, error) {

	err := r.layer.initialize()
	if err != nil {
		return Entity{}, err
	}

	if value {
		return newEntity(bitOneNode, r.storage), nil
	} else {
		return newEntity(bitZeroNode, r.storage), nil
	}
}
