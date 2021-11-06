package phrase

import (
	"errors"
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/concrete/bit"
	"github.com/mem-memov/semnet/internal/concrete/character"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"github.com/mem-memov/semnet/internal/concrete/word"
	"testing"
)

func TestRepository(t *testing.T) {
	data := []struct {
		name       string
		phrase     string
		provideErr error
		extractErr error
	}{
		{"no word", "", errors.New("no character in word cluster: "), nil},
		{"one word", "hello", nil, nil},
		{"many words", "this practical guide", nil, nil},
		{"short words", "a b c", nil, nil},
		{"words with the same beginning", "cat cab", nil, nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			slices := clew.NewSliceStorage()
			storage := clew.NewGraph(slices)
			classRepository := class.NewRepository(storage)
			bitRepository := bit.NewRepository(storage, classRepository)
			characterRepository := character.NewRepository(storage, classRepository, bitRepository)
			wordRepository := word.NewRepository(storage, classRepository, characterRepository)
			phraseRepository := NewRepository(storage, classRepository, wordRepository)

			aggregate, err := phraseRepository.Provide(d.phrase)
			if err != nil && err.Error() != d.provideErr.Error() {
				t.Fail()
			} else {
				return
			}

			phrase, err := aggregate.Extract()
			if err != nil && err.Error() != d.extractErr.Error() {
				t.Fail()
			} else {
				return
			}

			if phrase != d.phrase {
				t.Error(slices)
			}
		})
	}
}
