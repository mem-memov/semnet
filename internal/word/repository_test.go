package word

import (
	"errors"
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/class"
	"testing"
)

func TestRepository(t *testing.T) {
	data := []struct {
		name       string
		word       string
		provideErr error
		extractErr error
	}{
		{"no character", "", errors.New("no character in word entity: "), nil},
		{"one character", "a", nil, nil},
		{"many letters", "house", nil, nil},
		{"repeted letter", "pppp", nil, nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			slices := clew.NewSliceStorage()
			storage := clew.NewGraph(slices)
			classRepository := class.NewRepository(storage)
			bitRepository := bit.NewRepository(storage, classRepository)
			characterRepository := character.NewRepository(storage, classRepository, bitRepository)
			r := NewRepository(storage, classRepository, characterRepository)

			entity, err := r.Provide(d.word)
			if err != nil && err.Error() != d.provideErr.Error() {
				t.Fail()
			} else {
				return
			}

			word, err := r.Extract(entity)
			if err != nil && err.Error() != d.extractErr.Error() {
				t.Fail()
			} else {
				return
			}

			if word != d.word {
				t.Error(slices)
			}
		})
	}
}
