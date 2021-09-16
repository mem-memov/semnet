package detail

import (
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/bit"
	"github.com/mem-memov/semnet/internal/character"
	"github.com/mem-memov/semnet/internal/class"
	"github.com/mem-memov/semnet/internal/phrase"
	"github.com/mem-memov/semnet/internal/word"
	"testing"
)

func TestRepository(t *testing.T) {
	data := []struct {
		name       string
		object     string
		property   string
		provideErr error
		extractErr error
	}{
		{"no problem", "can think", "smart people", nil, nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			slices := clew.NewSliceStorage()
			storage := clew.NewGraph(slices)
			classRepository := class.NewRepository(storage)
			bitRepository := bit.NewRepository(storage, classRepository)
			characterRepository := character.NewRepository(storage, classRepository, bitRepository)
			wordRepository := word.NewRepository(storage, characterRepository)
			phraseRepository := phrase.NewRepository(storage, wordRepository)
			r := NewRepository(storage, phraseRepository)

			entity, err := r.Provide(d.object, d.property)
			if err != nil && err.Error() != d.provideErr.Error() {
				t.Fail()
			} else {
				return
			}

			object, property, err := r.Extract(entity)
			if err != nil && err.Error() != d.extractErr.Error() {
				t.Fail()
			} else {
				return
			}

			if object != d.object || property != d.property {
				t.Error(slices)
			}
		})
	}
}
