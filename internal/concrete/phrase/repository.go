package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Repository struct {
	phraseStorage   abstractPhrase.Storage
	phraseFactory abstractPhrase.Factory
	classRepository abstractClass.Repository
	wordRepository  abstractWord.Repository
	paths           *paths
}

var _ abstractPhrase.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository, wordRepository abstractWord.Repository) *Repository {

	phraseStorage := NewStorage(storage)

	return &Repository{
		phraseStorage:   phraseStorage,
		phraseFactory: NewFactory(phraseStorage),
		classRepository: classRepository,
		wordRepository:  wordRepository,
		paths:           newPaths(),
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

	class, err := r.classRepository.ProvideEntity()
	if err != nil {
		return nil, err
	}

	// root
	phrase, err := r.phraseFactory.ProvideEntity(class, firstWord_)
	if err != nil {
		return nil, err
	}

	// branches
out:
	for _, wordValue := range path[1:] {

		targetPhraseIdentifiers, err := phrase.GetTargetPhrases()
		if err != nil {
			return nil, err
		}

		targetPhrases := make([]abstractPhrase.Entity, 0, len(targetPhraseIdentifiers))
		for _, targetPhraseIdentifier := range targetPhraseIdentifiers {

			targetPhrase, err := r.phraseStorage.ReadEntityByPhrase(targetPhraseIdentifier)
			if err != nil {
				return nil, err
			}

			targetPhrases = append(targetPhrases, targetPhrase)
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
				phrase = targetPhrase.(Entity) // TODO: remove cast after word refactoring
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

		newPhrase, err := r.phraseFactory.ProvideEntity(class, word_)
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

		phraseIdentifier, err := phrase.GetSourcePhrase()
		if err != nil {
			return "", err
		}

		phrase, err := r.phraseStorage.ReadEntityByPhrase(phraseIdentifier)
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

func (r *Repository) Fetch(detail uint) (abstractPhrase.Entity, error) {

	return r.phraseStorage.ReadEntityByDetail(detail)
}
