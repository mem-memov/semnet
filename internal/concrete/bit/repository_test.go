package bit

import (
	"github.com/mem-memov/clew"
	"github.com/mem-memov/semnet/internal/concrete/class"
	"testing"
)

func TestRepository_ProvideZero(t *testing.T) {
	slices := clew.NewSliceStorage()
	storage := clew.NewGraph(slices)
	classRepository := class.NewRepository(storage)
	r := NewRepository(storage, classRepository)

	zeroEntity, err := r.Provide(false)
	if err != nil {
		t.Fail()
	}

	if zeroEntity.Identifier() != 13 {
		t.Errorf("wrong identifier %d", zeroEntity.Identifier())
	}

	// repeat

	zeroEntity, err = r.Provide(false)
	if err != nil {
		t.Fail()
	}

	if zeroEntity.Identifier() != 13 {
		t.Errorf("wrong identifier %d", zeroEntity.Identifier())
	}
}

func TestRepository_ProvideOne(t *testing.T) {
	slices := clew.NewSliceStorage()
	storage := clew.NewGraph(slices)
	classRepository := class.NewRepository(storage)
	r := NewRepository(storage, classRepository)

	oneEntity, err := r.Provide(true)
	if err != nil {
		t.Fail()
	}

	if oneEntity.Identifier() != 14 {
		t.Errorf("wrong identifier %d", oneEntity.Identifier())
	}

	// repeat

	oneEntity, err = r.Provide(true)
	if err != nil {
		t.Fail()
	}

	if oneEntity.Identifier() != 14 {
		t.Errorf("wrong identifier %d", oneEntity.Identifier())
	}
}
