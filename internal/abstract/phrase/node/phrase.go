package node

type Phrase interface {
	Identifier() uint
	NewPhrase(word Word) (Phrase, error)
	ReadTargets() ([]Phrase, error)
	ReadSources() ([]Phrase, error)
	GetClassAndWordAndDetail() (uint, uint, uint, error)
}
