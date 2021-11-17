package remark

type Repository interface {
	CreateFirstUserRemark(object string, property string) (Aggregate, error)
	GetRemark(remarkIdentifier uint) (Aggregate, error)
}
