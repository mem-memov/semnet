package word

import (
	"errors"
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"testing"
)

func TestRepository(t *testing.T) {
	data := []struct {
		name       string
		word       string
		provideErr error
		extractErr error
	}{
		{"no character", "", errors.New("no phrases in word entity: "), nil},
		{"one character", "a", nil, nil},
		{"many letters", "house", nil, nil},
		{"repeted letter", "pppp", nil, nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			slices := clew.NewSliceStorage()
			storage := clew.NewGraph(slices)
			bitRepository := bit.NewRepository(storage)
			characterRepository := character.NewRepository(storage, bitRepository)
			r := NewRepository(storage, characterRepository)

			entity, err := r.Provide(d.word)
			if err != nil && err.Error() != d.provideErr.Error() {
				t.Fail()
			} else {
				return
			}

			integer, err := r.Extract(entity)
			if err != nil && err.Error() != d.extractErr.Error() {
				t.Fail()
			} else {
				return
			}

			if integer != d.word {
				t.Error(slices)
			}
		})
	}
}
