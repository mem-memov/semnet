package character

type Elements interface {
	ProvideFirstElement(bitValue bool) (Element, error)
	CreateLastElement(character Entity) Element
}
