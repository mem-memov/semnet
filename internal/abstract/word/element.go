package word

type Element interface {
	GetEntity() Entity
	ProvideNextElement(characterValue rune) (Element, error)
	ExtractCharacterValue() (rune, error)
	HasPreviousElement() (bool, error)
	GetPreviousElement() (Element, error)
}
