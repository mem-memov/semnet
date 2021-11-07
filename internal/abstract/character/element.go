package character

type Element interface {
	GetEntity() Entity
	ProvideNextElement(bitValue bool) (Element, error)
	ExtractBitValue() (bool, error)
	HasPreviousElement() (bool, error)
	GetPreviousElement() (Element, error)
}
