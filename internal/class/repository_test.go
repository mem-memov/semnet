package class

import (
	"github.com/mem-memov/clew"
	"testing"
)

func TestRepository_ProvideEntity(t *testing.T) {
	slices := clew.NewSliceStorage()
	storage := clew.NewGraph(slices)
	r := NewRepository(storage)

	_, err := r.ProvideEntity()
	if err != nil {
		t.Fail()
	}

	_, err = r.ProvideEntity()
	if err != nil {
		t.Fail()
	}

	if result, err := storage.Has(9); err != nil || !result {
		t.Fail()
	}
}
