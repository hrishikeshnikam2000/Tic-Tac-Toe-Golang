package components

import (
	"testing"
)

//error remaining
func TestNewBoard(t *testing.T) {
	testBoard := []struct {
		size      uint8
		BoardSize uint8
		mark      string
	}{
		{3, 9, Nomark},
		{4, 16, Nomark},
	}

	for _, test := range testBoard {
		newBoard := NewBoard(test.size)
		gotCell := newBoard.Cells
		gotBoardsize := uint8(len(gotCell))
		if test.BoardSize != gotBoardsize {
			t.Error(test.BoardSize, gotBoardsize, " BoardSize Doesnot match")
		}

		for _, cell := range gotCell {
			gotMark := cell.GetMark()
			if test.mark != gotMark {
				t.Error(test.mark, gotMark, " The cell has Nomark(-) when it is created !")
			}

		}
	}
}

//error line 12
//{5, 24, Nomark},
//{2, 4, Xmark},
//error line 28
//fmt.Println(gotMark)
//
