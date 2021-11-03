package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Entity struct {
	class   uint
	word    uint
	phrase  uint
	detail  uint
	storage abstract.Storage
}

var _ abstractPhrase.Entity = Entity{}

func createEntity(storage abstract.Storage, classEntity abstractClass.Entity, wordEntity abstractWord.Entity) (Entity, error) {

	wordIdentifier, err := wordEntity.ProvideSingleTarget()
	if err != nil {
		return Entity{}, err
	}

	wordTargets, err := storage.ReadTargets(wordIdentifier)
	if err != nil {
		return Entity{}, err
	}

	switch len(wordTargets) {

	case 0:

		err = wordEntity.Mark(wordIdentifier)
		if err != nil {
			return Entity{}, err
		}

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

		err = storage.SetReference(class, word)
		if err != nil {
			return Entity{}, err
		}

		err = storage.SetReference(word, phrase)
		if err != nil {
			return Entity{}, err
		}

		err = storage.SetReference(phrase, detail)
		if err != nil {
			return Entity{}, err
		}

		err = wordEntity.PointToPhrase(word)
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

	case 1:

		if wordTargets[0] != wordEntity.PhraseIdentifier() {
			return Entity{}, fmt.Errorf("wrong target %d in detail tree at word %d", wordTargets[0], wordIdentifier)
		}

		return readEntityByWord(storage, wordIdentifier)

	default:
		return Entity{}, fmt.Errorf("wrong number of targets %d in word tree at word %d", len(wordTargets), wordIdentifier)
	}
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

func (e Entity) GetSourcePhrase() (abstractPhrase.Entity, error) {

	sourcePhrases, err := e.storage.ReadSources(e.phrase)
	if err != nil {
		return nil, err
	}

	if len(sourcePhrases) != 1 {
		return nil, fmt.Errorf("too many sources in phrase tree: %d at %d", len(sourcePhrases), e.phrase)
	}

	return readEntityByPhrase(e.storage, sourcePhrases[0])
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
