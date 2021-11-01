package phrase

import (
	"github.com/mem-memov/semnet/internal/abstract"
	abstractClass "github.com/mem-memov/semnet/internal/abstract/class"
	abstractPhrase "github.com/mem-memov/semnet/internal/abstract/phrase"
	abstractWord "github.com/mem-memov/semnet/internal/abstract/word"
)

type Repository struct {
	storage         abstract.Storage
	classRepository abstractClass.Repository
	wordRepository  abstractWord.Repository

	tree  abstractPhrase.Tree
	paths *paths
}

var _ abstractPhrase.Repository = &Repository{}

func NewRepository(storage abstract.Storage, classRepository abstractClass.Repository, wordRepository abstractWord.Repository) *Repository {

	return &Repository{
		storage:         storage,
		classRepository: classRepository,
		wordRepository:  wordRepository,

		tree:  newTree(storage, classRepository),
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

	phrase, err := r.tree.ProvideRoot(firstWord)
	if err != nil {
		return nil, err
	}

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

		newPhrase, err := createEntity(r.storage, class)
		if err != nil {
			return nil, err
		}

		err = word.PointToPhrase(newPhrase.GetWord())
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

	word, err := r.wordRepository.Fetch(phrase.GetTarget())
	if err != nil {
		return "", err
	}

	wordValue, err := word.
	if err != nil {
		return "", err
	}

	path := r.paths.create(wordValue)

	for {
		var isRoot bool
		phrase, isRoot, err = phrase.FindPrevious(r.entities)

		if isRoot {
			break
		}

		wordValue, err = phrase.WordValue()
		if err != nil {
			return "", err
		}

		path = append(path, wordValue)
	}

	return path.reverse().toString(), nil
}

func (r *Repository) Fetch(detailIdentifier uint) (abstractPhrase.Entity, error) {

	return r.entities.CreateWithDetail(detailIdentifier)
}
