package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractNode "github.com/mem-memov/semnet/internal/abstract/phrase/node"
)

type Entity struct {
	class   uint
	word    uint
	phrase  uint
	detail  uint
	storage abstract.Storage
}

var _ abstractPhrase.Entity = Entity{}

func createEntity(storage abstract.Storage, classEntity abstractClass.Entity) (Entity, error) {

	class, err := classEntity.CreatePhrase()
	if err != nil {
		return Entity{}, err
	}

	word, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	phrase, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	detail, err := storage.Create()
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(class, phrase)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(phrase, detail)
	if err != nil {
		return Entity{}, err
	}

	err = storage.SetReference(class, word)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: storage,
	}, nil
}

func readEntityByClass(storage abstract.Storage, class uint) (Entity, error) {

	_, word, err := storage.GetReference(class)
	if err != nil {
		return Entity{}, err
	}

	_, phrase, err := storage.GetReference(word)
	if err != nil {
		return Entity{}, err
	}

	_, detail, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: storage,
	}, nil
}

func readEntityByWord(storage abstract.Storage, word uint) (Entity, error) {

	class, phrase, err := storage.GetReference(word)
	if err != nil {
		return Entity{}, err
	}

	_, detail, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: storage,
	}, nil
}

func readEntityByPhrase(storage abstract.Storage, phrase uint) (Entity, error) {

	word, detail, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(word)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: storage,
	}, nil
}

func readEntityByDetail(storage abstract.Storage, detail uint) (Entity, error) {

	phrase, _, err := storage.GetReference(detail)
	if err != nil {
		return Entity{}, err
	}

	word, _, err := storage.GetReference(phrase)
	if err != nil {
		return Entity{}, err
	}

	class, _, err := storage.GetReference(word)
	if err != nil {
		return Entity{}, err
	}

	return Entity{
		class:   class,
		word:    word,
		phrase:  phrase,
		detail:  detail,
		storage: storage,
	}, nil
}

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

func (e Entity) GetTargetPhrases() ([]abstractPhrase.Entity, error) {

	targetPhraseIdentifiers, err := e.storage.ReadTargets(e.phrase)
	if err != nil {
		return nil, err
	}

	var targetPhrases []abstractPhrase.Entity

	for _, targetPhraseIdentifier := range targetPhraseIdentifiers {

		targetPhrase, err := readEntityByPhrase(e.storage, targetPhraseIdentifier)
		if err != nil {
			return nil, err
		}

		targetPhrases = append(targetPhrases, targetPhrase)
	}

	return targetPhrases, nil
}

func (e Entity) GetTargetWord() (uint, error) {

}

func (e Entity) DetailIdentifier() uint {
	return e.detailNode.Identifier()
}

func (e Entity) GetDetailNode() abstractNode.Detail {
	return e.detailNode
}

func (e Entity) ProvideDetailTarget(another abstractPhrase.Entity) (uint, error) {

	return e.detailNode.ProvideDetailTarget(another.GetDetailNode())
}

func (e Entity) Mark(sourceIdentifier uint) error {
	return e.detailNode.Mark(sourceIdentifier)
}

func (e Entity) HasWordValue(wordValue string) (bool, error) {

	hasWordValue, err := e.wordNode.HasWordValue(wordValue)
	if err != nil {
		return false, err
	}

	return hasWordValue, nil
}

func (e Entity) WordValue() (string, error) {

	return e.wordNode.WordValue()
}

func (e Entity) FindPrevious(entities abstractPhrase.Entities) (abstractPhrase.Entity, bool, error) {

	sourcePhrases, err := e.phraseNode.ReadSources()
	if err != nil {
		return Entity{}, false, nil
	}

	switch len(sourcePhrases) {
	case 0:
		return e, true, nil
	case 1:
		parentPhrase := sourcePhrases[0]

		classIdentifier, wordIdentifier, detailIdentifier, err := parentPhrase.GetClassAndWordAndDetail()
		if err != nil {
			return nil, false, nil
		}

		return entities.Create(classIdentifier, wordIdentifier, parentPhrase.Identifier(), detailIdentifier), false, nil
	default:
		return nil, false, fmt.Errorf("too many sources in phrase tree")
	}
}
