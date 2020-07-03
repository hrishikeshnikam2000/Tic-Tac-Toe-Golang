package services

import (
	"fmt"
	"tic_tac_toe/components"
)

type BoardService struct {
	Board *components.Board
}

func NewBoardService(board *components.Board) *BoardService {
	newBoardService := &BoardService{
		Board: board,
	}
	return newBoardService
}

func (b *BoardService) PutMarkInPosition(player *components.Player, position uint8) error {
	err := b.Board.Cells[position].SetMark(player.Mark)
	return err
}

func (b *BoardService) PrintBoard() string {
	boardSize := b.Board.Size * b.Board.Size
	boardStr := ""
	var i uint8
	for i = 0; i < boardSize; i++ {
		if (i % b.Board.Size) < (b.Board.Size - 1) {
			boardStr += fmt.Sprintf("_%s_|", b.Board.Cells[i].GetMark())
		}
		if (i % b.Board.Size) == (b.Board.Size - 1) {
			boardStr += fmt.Sprintf("_%s_\n", b.Board.Cells[i].GetMark())
		}
	}
	return boardStr
}

func (b *BoardService) CheckBoardIsFull() bool {
	for _, cell := range b.Board.Cells {
		if cell.GetMark() == components.Nomark {
			return false
		}
	}
	return true
}
