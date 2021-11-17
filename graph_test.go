package semnet

import (
	"github.com/mem-memov/clew"
	"testing"
)

func TestNewGraph(t *testing.T) {
	slices := clew.NewSliceStorage()
	storage := clew.NewGraph(slices)
	graph := NewGraph(storage)
	output := NewOutput(graph)

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

	goesToPlay, err := goesACat.AddRemarkToFact("to play")
	if err != nil {
		t.Fail()
	}

	goes, toPlay, err := goesToPlay.GetObjectAndProperty()
	if err != nil {
		t.Fail()
	}

	if goes != "goes" {
		t.Fail()
	}

	if toPlay != "to play" {
		t.Fail()
	}

	hidesTheMouse, err := goesACat.AddFactToStory("hides", "the mouse")
	if err != nil {
		t.Fail()
	}

	hides, theMouse, err := hidesTheMouse.GetObjectAndProperty()
	if err != nil {
		t.Fail()
	}

	if hides != "hides" {
		t.Fail()
	}

	if theMouse != "the mouse" {
		t.Fail()
	}

	remark, err := graph.GetRemark(hidesTheMouse.GetIdentifier())

	story, err := output.GetStory(remark)
	if err != nil {
		t.Fail()
	}

	if story != "" {
		t.Errorf(story)
	}
}
