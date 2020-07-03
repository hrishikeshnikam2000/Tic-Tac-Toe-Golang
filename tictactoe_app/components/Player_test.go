package components

import "testing"

func TestNewPlayer(t *testing.T) {
	actualPlayer := NewPlayer("Fulva", "X")
	expectedPlayer := &Player{Name: "Fulva", Mark: "X"}
	if actualPlayer.Name != expectedPlayer.Name {
		t.Error("Got Name is", actualPlayer.Name, "but,Want Name is:", expectedPlayer.Name)
	}
	if actualPlayer.Mark != expectedPlayer.Mark {
		t.Error("Got Mark is", actualPlayer.Mark, "but,Want Mark is:", expectedPlayer.Mark)
	}
}
