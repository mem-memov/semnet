package semnet

import (
	"github.com/mem-memov/clew"
	"testing"
)

func TestNewGraph(t *testing.T) {
	slices := clew.NewSliceStorage()
	storage := clew.NewGraph(slices)
	graph := NewGraph(storage)

	goesACat, err := graph.CreateStory("goes", "a cat")
	if err != nil {
		t.Fail()
	}

	goes, aCat, err := goesACat.GetObjectAndProperty()
	if err != nil {
		t.Fail()
	}

	if goes != "goes" {
		t.Fail()
	}

	if aCat != "a cat" {
		t.Fail()
	}
}
