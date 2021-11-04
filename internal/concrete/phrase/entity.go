package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
)

type Entity struct {
	class   uint
	word    uint
	phrase  uint
	detail  uint
	storage abstract.Storage
}

var _ abstractPhrase.Entity = Entity{}

func (e Entity) GetClass() uint {

	return e.class
}

func (e Entity) GetWord() uint {

	return e.word
}

func (e Entity) GetPhrase() uint {

	return e.phrase
}

func (e Entity) GetDetail() uint {

	return e.detail
}

func (e Entity) PointToPhrase(phrase uint) error {

	return e.storage.Connect(e.phrase, phrase)
}

func (e Entity) GetTargetPhrases() ([]uint, error) {

	return e.storage.ReadTargets(e.phrase)
}

func (e Entity) GetSourceWord() (uint, error) {

	sourceWords, err := e.storage.ReadSources(e.word)
	if err != nil {
		return 0, err
	}

	if len(sourceWords) != 1 {
		return 0, fmt.Errorf("phrase has wrong number of source words: %d at %d", len(sourceWords), e.word)
	}

	return sourceWords[0], nil
}

func (e Entity) HasSourcePhrase() (bool, error) {

	// TODO: read count from DB
	sourcePhrases, err := e.storage.ReadSources(e.phrase)
	if err != nil {
		return false, err
	}

	return len(sourcePhrases) != 0, nil
}

func (e Entity) GetSourcePhrase() (uint, error) {

	sourcePhrases, err := e.storage.ReadSources(e.phrase)
	if err != nil {
		return 0, err
	}

	if len(sourcePhrases) != 1 {
		return 0, fmt.Errorf("too many sources in phrase tree: %d at %d", len(sourcePhrases), e.phrase)
	}

	return sourcePhrases[0], nil
}

func (e Entity) GetSourceDetails() ([]uint, error) {

	return e.storage.ReadSources(e.detail)
}

func (e Entity) GetTargetDetails() ([]uint, error) {

	return e.storage.ReadTargets(e.detail)
}

func (e Entity) AddSourceDetail(detail uint) error {

	return e.storage.Connect(detail, e.detail)
}

func (e Entity) AddTargetDetail(detail uint) error {

	return e.storage.Connect(e.detail, detail)
}

func (e Entity) Mark(sourceIdentifier uint) error {

	return e.storage.Connect(sourceIdentifier, e.detail)
}
