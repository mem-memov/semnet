package code

import (
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/bit"
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
			bitRepository := bit.NewRepository(storage)
			r := newRepository(storage, bitRepository)

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
