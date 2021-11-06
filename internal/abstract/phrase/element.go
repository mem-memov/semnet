package phrase

type Element interface {
	GetEntity() Entity
	ProvideNextElement(wordValue string) (Element, error)
	ExtractWordValue() (string, error)
	HasPreviousElement() (bool, error)
	GetPreviousElement() (Element, error)
}
