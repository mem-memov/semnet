package action

//import (
//	"github.com/mem-memov/semnet/internal/word"
//	"strings"
//)
//
//type actions struct {
//	storage storage
//	word   *word
//}
//
//func newActions(storage storage, word *word) *actions {
//	return &actions{
//		storage: storage,
//		word:   word,
//	}
//}
//
//func (a *actions) create(name string) (Action, error) {
//	wordNames := strings.Split(name, " ")
//
//	words := make([]word.Word, len(wordNames))
//
//	for i, wordName := range wordNames {
//		word, err := a.word.create(wordName)
//		if err != nil {
//			return Action{}, err
//		}
//		words[i] = word
//	}
//}
