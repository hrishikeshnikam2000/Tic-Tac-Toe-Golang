package services

import (
	"errors"
	"testing"
	"tic_tac_toe/components"
)

func TestNewBoardService(t *testing.T) {
	wantBoardService := &BoardService{Board: components.NewBoard(2)}
	got := NewBoardService(components.NewBoard(2))
	var i uint8
	for i = 0; i < got.Board.Size*got.Board.Size; i++ {
		if wantBoardService.Board.Cells[i].Mark != got.Board.Cells[i].Mark {
			t.Error("The cell has Nomark when it is created !")
		}
	}
	//fmt.Println(wantBoardService.PrintBoard())
	//fmt.Println(got.PrintBoard())
}

func TestPutMarkInPosition(t *testing.T) {
	testPutMarks := []struct {
		wantBoardService *BoardService
		wantPlayer       *components.Player
		wantPosition     uint8
		wantError        error
	}{
		{&BoardService{Board: components.NewBoard(2)}, components.NewPlayer("Fulva", components.Xmark), 3, nil},
		{&BoardService{Board: components.NewBoard(2)}, components.NewPlayer("XYZ", components.Xmark), 2, nil},
		{&BoardService{
			Board: &components.Board{
				Cells: []*components.Cell{
					&components.Cell{Mark: components.Xmark},
					&components.Cell{Mark: components.Nomark},
					&components.Cell{Mark: components.Omark},
					&components.Cell{Mark: components.Xmark},
				},
				Size: 2,
			},
		}, &components.Player{Name: "Sonu", Mark: components.Omark},
			//1, gotError is nil because cell[1] has Nomark
			3, //gotError is !=nil because cell[3] is already marked !
			errors.New("Cell is already marked !"),
		},
	}
	for _, test := range testPutMarks {
		gotError := test.wantBoardService.PutMarkInPosition(test.wantPlayer, test.wantPosition)
		//fmt.Println(gotError) Output:<nil> <nil> The cell is already marked !
		//fmt.Println(test.wantBoardService.PrintBoard())
		if gotError != nil || test.wantError != nil {
			t.Error(gotError, test.wantError)
		}
	}
}

func TestCheckBoardIsFull(t *testing.T) {
	testFull := []struct {
		wantBoardService *BoardService
		wantResult       bool
	}{
		{&BoardService{Board: components.NewBoard(2)}, false},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Omark},
			},
			Size: 2,
		},
		},
			false},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Omark},
			},
			Size: 2,
		},
		},
			true},
	}
	for _, test := range testFull {
		gotResult := test.wantBoardService.CheckBoardIsFull()
		//fmt.Println(gotResult) Output: false false true
		//fmt.Println(test.wantResult) Output: false false true
		if gotResult != test.wantResult {
			t.Error(test.wantResult, gotResult)
		}
	}
}
