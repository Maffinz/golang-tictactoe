package service

import (
	"errors"
	"fmt"
)

type Game struct {
	Board         *Board
	NumberOfTurns int
	GameOver      bool
	CurrentPlayer string
}

func InitializeGame() *Game {
	return &Game{
		Board:         initializeBoard(),
		NumberOfTurns: 0,
		GameOver:      false,
	}
}

func (g *Game) GameLoop() {
	WelcomePrompt()

	for !g.GameOver {
		g.GameLogic()
	}

	g.GameOverPrompt(g.CurrentPlayer)
}

func WelcomePrompt() {
	fmt.Println("Welcome to Tic-Tac-Toe")
}

func (g *Game) GameLogic() {
	g.Board.PrintBoard()

	g.CurrentPlayer = ChoosePlayer(g.NumberOfTurns)

	var loc int
	fmt.Print("enter location: ")
	fmt.Scanln(&loc)

	err := g.Board.PlaceOnBoard(g.CurrentPlayer, ChooseLocation(g.CurrentPlayer, loc))
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	} else {
		g.IncrementGameTurns()
	}

	g.GameOver = g.Board.CheckBoardForWinner()
}

func (g *Game) IncrementGameTurns() {
	g.NumberOfTurns++
}

func (g *Game) PlayerSelection(player string, loc int) error {
	err := g.Board.PlaceOnBoard(player, ChooseLocation(player, loc))
	if err != nil {
		return errors.New("something happend when placing on board")
	}

	g.IncrementGameTurns()
	return nil
}

func (g *Game) GameOverPrompt(player string) {
	fmt.Printf("Game Over, player %s won!\n", player)
	fmt.Printf("Number of turns: %d\n\n", g.NumberOfTurns)

	g.Board.PrintBoard()
}
