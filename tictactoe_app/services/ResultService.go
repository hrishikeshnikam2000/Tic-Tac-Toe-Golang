package services

import (
	"tic_tac_toe/components"
)

type ResultService struct {
	BoardService *BoardService
}

type Result struct {
	Win  bool
	Draw bool
}

func NewResultService(boardService *BoardService) *ResultService {
	resultService := &ResultService{
		BoardService: boardService,
	}
	return resultService
}

func (rs *ResultService) checkRows(mark string, position uint8) bool {
	size := rs.BoardService.Board.Size
	start := size * (position / size)
	for i := start; i < start+size; i++ {
		if rs.BoardService.Board.Cells[i].GetMark() != mark {
			return false
		}
	}
	return true
}

func (rs *ResultService) checkColumns(mark string, position uint8) bool {
	size := rs.BoardService.Board.Size
	start := position % size
	for i := start; i < size*size; i += size {
		if rs.BoardService.Board.Cells[i].GetMark() != mark {
			return false
		}
	}
	return true
}

func (rs *ResultService) checkLeftToRightDiagonal(mark string, position uint8) bool {
	size := rs.BoardService.Board.Size
	start := (position % size) - (position / size)
	if start != 0 {
		return false
	}
	for i := start; i < size*size; i += (size + 1) {
		if rs.BoardService.Board.Cells[i].GetMark() != mark {
			return false
		}
	}
	return true
}

func (rs *ResultService) checkRightToLeftDiagonal(mark string, position uint8) bool {
	size := rs.BoardService.Board.Size
	start := (position % size) + (position / size)
	if start != (size - 1) {
		return false
	}
	for i := start; i <= size*(size-1); i += (size - 1) {
		if rs.BoardService.Board.Cells[i].GetMark() != mark {
			return false
		}
	}
	return true
}

func (rs *ResultService) GetResult(player *components.Player, position uint8) Result {
	if rs.checkRows(player.Mark, position) {
		return Result{true, false}
	} else if rs.checkColumns(player.Mark, position) {
		return Result{true, false}
	} else if rs.checkLeftToRightDiagonal(player.Mark, position) {
		return Result{true, false}
	} else if rs.checkRightToLeftDiagonal(player.Mark, position) {
		return Result{true, false}
	} else if rs.BoardService.CheckBoardIsFull() {
		return Result{false, true}
	}
	return Result{false, false}

}
