package character

import (
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/concrete/bit"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"testing"
)

func TestRepository(t *testing.T) {
	data := []struct {
		name       string
		integer    int32
		provideErr error
		extractErr error
	}{
		{"0", 0, nil, nil},
		{"1", 1, nil, nil},
		{"10", 2, nil, nil},
		{"11", 3, nil, nil},
		{"111", 7, nil, nil},
		{"1000", 8, nil, nil},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			slices := clew.NewSliceStorage()
			storage := clew.NewGraph(slices)
			classRepository := class.NewRepository(storage)
			bitRepository := bit.NewRepository(storage, classRepository)
			r := NewRepository(storage, classRepository, bitRepository)

			entity, err := r.Provide(d.integer)

			if err != d.provideErr {
				t.Fail()
			}

			integer, err := r.Extract(entity)

			if err != d.extractErr {
				t.Fail()
			}

			if integer != d.integer {
				t.Error(entity, slices)
			}
		})
	}
}
