package services

import (
	"errors"
	"tic_tac_toe/components"
)

type GameService struct {
	ResultService *ResultService
	Players       [2]*components.Player
	Turn          int
}

func NewGameService(resultService *ResultService, players [2]*components.Player) *GameService {
	gameService := &GameService{
		ResultService: resultService,
		Players:       players,
		Turn:          0,
	}
	return gameService
}

func (gs *GameService) Play(position uint8) (Result, error) {
	size := gs.ResultService.BoardService.Board.Size
	if position < 0 || position >= size*size {
		return Result{false, false}, errors.New("position is not valid !")
	}
	var result Result
	if gs.Turn%2 == 0 {
		err := gs.ResultService.BoardService.PutMarkInPosition(gs.Players[0], position)
		if err != nil {
			return Result{false, false}, err
		}
		result = gs.ResultService.GetResult(gs.Players[0], position)
	} else if gs.Turn%2 == 1 {
		err := gs.ResultService.BoardService.PutMarkInPosition(gs.Players[1], position)
		if err != nil {
			return Result{false, false}, err
		}
		result = gs.ResultService.GetResult(gs.Players[1], position)
	}
	gs.Turn++
	return result, nil
}

/* func checkError(err error)  {
	if err != nil {
		return Result{false, false}, err
	}
}
*/
