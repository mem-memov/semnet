package detail

type Repository interface {
	Provide(object string, property string) (Aggregate, error)
	Fetch(remark uint) (Aggregate, error)
}
