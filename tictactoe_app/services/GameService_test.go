package services

import (
	"errors"
	"testing"
	"tic_tac_toe/components"
)

func TestPlay(t *testing.T) {
	testsPlay := []struct {
		wantGameService *GameService
		wantPosition    uint8
		wantError       error
		wantResult      Result
	}{
		{&GameService{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Omark},
			},
			Size: 3,
		}}}, [2]*components.Player{
			{Name: "Fulva", Mark: components.Xmark},
			{Name: "Sonu", Mark: components.Omark},
		}, 6,
		}, 3, nil, Result{true, false}},

		{&GameService{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Omark},
			},
			Size: 3,
		}}}, [2]*components.Player{
			{Name: "Fulva", Mark: components.Xmark},
			{Name: "Sonu", Mark: components.Omark},
		}, 6,
		}, 4, errors.New("already marked!"), Result{false, false}},

		{&GameService{&ResultService{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Omark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Xmark},
				&components.Cell{Mark: components.Nomark},
				&components.Cell{Mark: components.Omark},
			},
			Size: 3,
		}}}, [2]*components.Player{
			{Name: "Fulva", Mark: components.Xmark},
			{Name: "Sonu", Mark: components.Omark},
		}, 7,
		}, 9, errors.New("Position not valid !"), Result{false, false}},
	}
	for _, test := range testsPlay {
		gotResult, goterror := test.wantGameService.Play(test.wantPosition)
		//fmt.Println(goterror)
		if goterror != nil || test.wantError != nil {
			t.Error(goterror, test.wantError)
		}
		//fmt.Println(gotResult)
		if gotResult != test.wantResult {
			t.Error(gotResult, test.wantResult)
		}
	}

}
