package word

type Elements interface {
	ProvideFirstElement(characterValue rune) (Element, error)
	CreateLastElement(word Entity) Element
}
