package node

type Word interface {
	Identifier() uint
	NewWord(wordValue string) (Word, error)
	HasWordValue(value string) (bool, error)
	WordValue() (string, error)
}
