package phrase

import (
	"fmt"
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Repository struct {
	storage         abstract.Storage
	classRepository abstractClass.Repository
	wordRepository  abstractWord.Repository
	paths *paths
}

var _ abstractPhrase.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository, wordRepository abstractWord.Repository) *Repository {

	return &Repository{
		storage:         storage,
		classRepository: classRepository,
		wordRepository:  wordRepository,
		paths: newPaths(),
	}
}

func (r *Repository) Provide(words string) (abstractPhrase.Entity, error) {

	path, err := r.paths.collect(words)
	if err != nil {
		return nil, err
	}

	firstWord, err := r.wordRepository.Provide(path[0])
	if err != nil {
		return nil, err
	}

	// TODO: remove after refactoring word package
	firstWord_ := firstWord.(abstractWord.Entity)

	// root
	wordIdentifier, err := firstWord_.ProvideSingleTarget()
	if err != nil {
		return nil, err
	}

	wordTargets, err := r.storage.ReadTargets(wordIdentifier)
	if err != nil {
		return nil, err
	}

	var phrase abstractPhrase.Entity
	switch len(wordTargets) {
	case 0:
		err = firstWord_.Mark(wordIdentifier)
		if err != nil {
			return nil, err
		}

		class, err := r.classRepository.ProvideEntity()
		if err != nil {
			return nil, err
		}

		phrase, err = createEntity(r.storage, class)
		if err != nil {
			return nil, err
		}

	case 1:
		if wordTargets[0] != firstWord_.PhraseIdentifier() {
			return nil, fmt.Errorf("wrong target %d in detail tree at word %d", wordTargets[0], wordIdentifier)
		}

		phrase, err = readEntityByWord(r.storage, wordIdentifier)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("wrong number of targets %d in word tree at word %d", len(wordTargets), wordIdentifier)
	}

	// branches
out:
	for _, wordValue := range path[1:] {

		targetPhrases, err := phrase.GetTargetPhrases()
		if err != nil {
			return nil, err
		}

		// use existing
		for _, targetPhrase := range targetPhrases {
			targetWord, err := r.wordRepository.Fetch(targetPhrase.GetWord())
			if err != nil {
				return nil, err
			}

			targetValue, err := r.wordRepository.Extract(targetWord)
			if err != nil {
				return nil, err
			}

			if wordValue == targetValue {
				phrase = targetPhrase
				break out
			}
		}

		// create new
		class, err := r.classRepository.ProvideEntity()
		if err != nil {
			return nil, err
		}

		word, err := r.wordRepository.Provide(wordValue)
		if err != nil {
			return nil, err
		}

		// TODO: remove after refactoring word package
		word_ := word.(abstractWord.Entity)

		newPhrase, err := createEntity(r.storage, class)
		if err != nil {
			return nil, err
		}

		err = word_.PointToPhrase(newPhrase.GetWord())
		if err != nil {
			return nil, err
		}

		err = phrase.PointToPhrase(newPhrase.GetPhrase())
		if err != nil {
			return nil, err
		}

		phrase = newPhrase
	}

	return phrase, nil
}

func (r *Repository) Extract(phrase abstractPhrase.Entity) (string, error) {

	sourceWordIdentifier, err := phrase.GetSourceWord()
	if err != nil {
		return "", err
	}

	word, err := r.wordRepository.Fetch(sourceWordIdentifier)
	if err != nil {
		return "", err
	}

	wordValue, err := r.wordRepository.Extract(word)
	if err != nil {
		return "", err
	}


	path := r.paths.create(wordValue)

	for {
		hasSourcePhrase, err := phrase.HasSourcePhrase()
		if err != nil {
			return "", err
		}

		if !hasSourcePhrase {
			break
		}

		phrase, err = phrase.GetSourcePhrase()
		if err != nil {
			return "", err
		}

		sourceWordIdentifier, err := phrase.GetSourceWord()
		if err != nil {
			return "", err
		}

		word, err := r.wordRepository.Fetch(sourceWordIdentifier)
		if err != nil {
			return "", err
		}

		wordValue, err := r.wordRepository.Extract(word)
		if err != nil {
			return "", err
		}


		path = append(path, wordValue)
	}

	return path.reverse().toString(), nil
}

func (r *Repository) Fetch(detailIdentifier uint) (abstractPhrase.Entity, error) {

	return readEntityByDetail(r.storage, detailIdentifier)
}
