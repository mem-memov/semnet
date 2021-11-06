package word

type Element interface {
	ProvideNextElement(wordValue string) (Element, error)
}
