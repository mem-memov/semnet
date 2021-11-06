package phrase

type Elements interface {
	ProvideFirstElement(wordValue string) (Element, error)
	CreateLastElement(phrase Entity) Element
}
