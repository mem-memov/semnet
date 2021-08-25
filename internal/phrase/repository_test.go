package phrase

import (
	"errors"
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/word"
	"testing"
)

func TestRepository(t *testing.T) {
	data := []struct {
		name       string
		phrase     string
		provideErr error
		extractErr error
	}{
		{"no word", "", errors.New("no character in word entity: "), nil},
		{"one word", "hello", nil, nil},
		{"many words", "this practical guide", nil, nil},
		{"short words", "a b c", nil, nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			slices := clew.NewSliceStorage()
			storage := clew.NewGraph(slices)
			bitRepository := bit.NewRepository(storage)
			characterRepository := character.NewRepository(storage, bitRepository)
			wordRepository := word.NewRepository(storage, characterRepository)
			r := NewRepository(storage, wordRepository)

			entity, err := r.Provide(d.phrase)
			if err != nil && err.Error() != d.provideErr.Error() {
				t.Fail()
			} else {
				return
			}

			phrase, err := r.Extract(entity)
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
