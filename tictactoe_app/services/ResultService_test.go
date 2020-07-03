package services

import (
	"testing"
	"tic_tac_toe/components"
)

func TestCheckRows(t *testing.T) {
	testForRow := []struct {
		wantResultService *ResultService
		wantMark          string
		wantPosition      uint8
		wantResult        bool
	}{
		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
					},
					Size: 2,
				},
			},
		}, components.Xmark, 1, true,
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
					},
					Size: 3,
				},
			},
		}, components.Xmark, 6, true,
		},
	}
	for _, testRow := range testForRow {
		got := testRow.wantResultService.checkRows(testRow.wantMark, testRow.wantPosition)
		//fmt.Println(got)
		if got != testRow.wantResult {
			t.Error(got, testRow.wantResult)
		}
	}
}

func TestCheckColumns(t *testing.T) {
	testForColumn := []struct {
		wantResultService *ResultService
		wantMark          string
		wantPosition      uint8
		wantResult        bool
	}{
		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
					},
					Size: 2,
				},
			},
		}, components.Omark, 3, true,
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
					},
					Size: 3,
				},
			},
		}, components.Xmark, 4, true,
		},
	}
	for _, testColumn := range testForColumn {
		got := testColumn.wantResultService.checkColumns(testColumn.wantMark, testColumn.wantPosition)
		//fmt.Println(got)
		if got != testColumn.wantResult {
			t.Error(got, testColumn.wantResult)
		}
	}
}

func TestCheckLeftToRightDiagonal(t *testing.T) {
	testForLRDiagonal := []struct {
		wantResultService *ResultService
		wantMark          string
		wantPosition      uint8
		wantResult        bool
	}{
		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
					},
					Size: 2,
				},
			},
		}, components.Xmark, 3, true,
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Omark},
					},
					Size: 3,
				},
			},
		}, components.Omark, 4, true,
		},
	}
	for _, testLR := range testForLRDiagonal {
		got := testLR.wantResultService.checkLeftToRightDiagonal(testLR.wantMark, testLR.wantPosition)
		//fmt.Println(got)
		//fmt.Println(testLR.wantResult)
		if got != testLR.wantResult {
			t.Error(got, testLR.wantResult)
		}
	}
}

func TestRightToLeftDiagonal(t *testing.T) {
	testForRLDiagonal := []struct {
		wantResultService *ResultService
		wantMark          string
		wantPosition      uint8
		wantResult        bool
	}{
		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Nomark},
					},
					Size: 2,
				},
			},
		}, components.Xmark, 2, true,
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Nomark},
					},
					Size: 3,
				},
			},
		}, components.Omark, 4, true,
		},
	}
	for _, testRL := range testForRLDiagonal {
		got := testRL.wantResultService.checkRightToLeftDiagonal(testRL.wantMark, testRL.wantPosition)
		//fmt.Println(got)
		if got != testRL.wantResult {
			t.Error(got, testRL.wantResult)
		}
	}
}

func TestGetResult(t *testing.T) {
	testForResult := []struct {
		wantResultService *ResultService
		wantPlayer        *components.Player
		wantPosition      uint8
		wantResult        Result
	}{
		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Nomark},
					},
					Size: 2,
				},
			},
		}, &components.Player{Name: "Fulva", Mark: components.Xmark}, 2, Result{true, false},
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Xmark},
					},
					Size: 2,
				},
			},
		}, &components.Player{Name: "ABC", Mark: components.Xmark}, 3, Result{true, false},
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
					},
					Size: 2,
				},
			},
		}, &components.Player{Name: "XYZ", Mark: components.Omark}, 3, Result{true, false},
		},

		{&ResultService{
			&BoardService{
				&components.Board{
					Cells: []*components.Cell{
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Xmark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Omark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Nomark},
						&components.Cell{Mark: components.Nomark},
					},
					Size: 3,
				},
			},
		}, &components.Player{Name: "Sonu", Mark: components.Omark}, 5, Result{true, false},
		},
	}
	for _, testResult := range testForResult {
		got := testResult.wantResultService.GetResult(testResult.wantPlayer, testResult.wantPosition)
		//fmt.Println(got)
		if got != testResult.wantResult {
			t.Error(got, testResult.wantResult)
		}
	}
}
