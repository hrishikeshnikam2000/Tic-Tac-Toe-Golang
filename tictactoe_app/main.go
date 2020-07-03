package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tic_tac_toe/components"
	"tic_tac_toe/services"
)

//global variables
var position int
var Size int
var reader *bufio.Reader
var boardService *services.BoardService
var resultService *services.ResultService
var gameService *services.GameService
var err error

func main() {
	var board *components.Board
	fmt.Println("----|-x-o-|-----|-x-o-|-----TIC-TAC-TOE-----|-x-o-|-----|-x-o-|----")
	fmt.Println("**********Prefered board size either should be eighter 3x3 or 4x4 or the last option is 5x5**********")
	fmt.Println("Enter 3 for 3x3 , 4 for 4x4 and 5 for 5x5 board size")

	for {
		fmt.Print("Enter your desired board Size ---->")
		reader = bufio.NewReader(os.Stdin)
		size, err := reader.ReadString('\n')
		checkError(err)
		size = strings.TrimSpace(size)
		Size, err = strconv.Atoi(size)
		if err != nil {
			fmt.Println("Please enter an Integer value")
		} else if Size < 3 || Size > 5 {
			fmt.Println("Invalid size, Please select the value 3,4 or 5")
		} else {
			fmt.Printf("The Board size is %d*%d", Size, Size)
			fmt.Println()
			fmt.Println()
			board = components.NewBoard(uint8(Size))
			break
		}
	}

	fmt.Println("---*---|x-o|---*---Enter Information of Player1---*---|x-o|---*---")
	var player1 *components.Player
	var player1Name, player1Mark string
	var err error
	fmt.Println("Enter the Name of first player(first letter capital):")
	reader = bufio.NewReader(os.Stdin)
	player1Name, err = reader.ReadString('\n')
	checkError(err)
	player1Name = strings.TrimSpace(player1Name)

	for {
		fmt.Println("Enter Mark for Player1(X or O(enter in Capital))")
		reader = bufio.NewReader(os.Stdin)
		player1Mark, err = reader.ReadString('\n')
		checkError(err)
		player1Mark = strings.TrimSpace(player1Mark)
		if player1Mark != components.Xmark && player1Mark != components.Omark {
			fmt.Println("Invalid mark,Please select either X or O")
		} else {
			fmt.Printf("Player %s play with mark %s", player1Name, player1Mark)
			fmt.Println()
			fmt.Println()
			player1 = components.NewPlayer(player1Name, player1Mark)
			break
		}
	}

	fmt.Println("---*---|x-o|---*---Enter Information of Player2---*---|x-o|---*---")
	var player2 *components.Player
	var player2Name, player2Mark string
	for {
		fmt.Println("Enter the Name of Second player(first letter capital):")
		reader = bufio.NewReader(os.Stdin)
		player2Name, err = reader.ReadString('\n')
		checkError(err)
		player2Name = strings.TrimSpace(player2Name)
		if player2Name == player1Name {
			fmt.Println("This playerName is already taken")
		} else {
			break
		}
	}

	player2Mark = ""
	if player1Mark == components.Xmark {
		player2Mark = components.Omark
	} else if player1Mark == components.Omark {
		player2Mark = components.Xmark
	}

	fmt.Printf("Player %s play with mark %s", player2Name, player2Mark)
	fmt.Println()
	fmt.Println()
	player2 = components.NewPlayer(player2Name, player2Mark)

	boardService = services.NewBoardService(board)
	resultService = services.NewResultService(boardService)
	gameService = services.NewGameService(resultService, [2]*components.Player{player1, player2})

	fmt.Println(".........................The Game Begins!!!.....................................")
	fmt.Println("~~~  Good luck to both of you! May the better person win !!  ~~~   xoxo :) ")

	var i uint8
	for i = 0; i < uint8(Size*Size); i++ {
		fmt.Println(gameService.ResultService.BoardService.PrintBoard())
		if i%2 == 0 {
			res := TurnPlayer(player1.Name)
			if res == true {
				return
			}

		} else {
			res := TurnPlayer(player2.Name)
			if res == true {
				return
			}
		}
	}
}

func TurnPlayer(player string) bool {
reSelect:
	fmt.Print(player, ", Enter the position you want to place the mark:")
	reader = bufio.NewReader(os.Stdin)
	pos, err := reader.ReadString('\n')
	checkError(err)
	pos = strings.TrimSpace(pos)
	position, err = strconv.Atoi(pos)
	if err != nil {
		fmt.Println("Position should be an integer")
		goto reSelect
	}
	result, err := gameService.Play(uint8(position))
	if err != nil {
		fmt.Println(err)
		if err.Error() == "The cell has already been marked !" {
			goto reSelect
		}
		if err.Error() == "the position you have entered is not valid !" {
			goto reSelect
		}
	}

	if result.Win == true {
		fmt.Println(gameService.ResultService.BoardService.PrintBoard())
		fmt.Println(player, " you are the champion! ")
		return true
	}

	if result.Draw == true {
		fmt.Println(gameService.ResultService.BoardService.PrintBoard())
		fmt.Println("Game has been Drawn")
		return true
	}
	return false

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//imported services and components
//using bufio for taking input
