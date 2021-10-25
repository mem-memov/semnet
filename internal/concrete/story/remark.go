package story

type Remark interface {
	GetFact() (Fact, error)
	GetObjectProperty() (string, string, error)
}
