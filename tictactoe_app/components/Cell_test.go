package components

import (
	"fmt"
	"testing"
)

func TestNewCell(t *testing.T) {
	actualCellMark := NewCell()
	expectedCellMark := &Cell{Mark: "-"}

	if *actualCellMark != *expectedCellMark {
		t.Error("The cell has Nomark(-) when it is created !")
	}
}

func TestSetMark(t *testing.T) {
	var testMark = []struct {
		OldMark   string
		newMark   string
		wantError error
	}{
		{"-", "X", nil},
		{"-", "O", nil},
	}
	for _, mark := range testMark {
		cell := NewCell()
		actualError := cell.SetMark(mark.newMark)
		fmt.Println(cell.GetMark())
		if actualError != mark.wantError {
			t.Error("Got Error is :", actualError, "but, Want Error is :", mark.wantError)
		}
	}
}
