package phrase

type Aggregate interface {
	AddSourceDetail(detail uint) error
	AddTargetDetail(detail uint) error
	GetSourceDetails() ([]uint, error)
	GetTargetDetails() ([]uint, error)
	Extract() (string, error)
}
