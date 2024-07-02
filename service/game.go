package service

import "fmt"

type Game struct {
	board         *Board
	numberOfTurns int
	gameOver      bool
}

func InitializeGame() *Game {
	return &Game{
		board:         initializeBoard(),
		numberOfTurns: 0,
		gameOver:      false,
	}
}

func (g *Game) GameLoop() {
	WelcomePrompt()

	for !g.gameOver {
		g.GameLogic()
	}
}

func WelcomePrompt() {
	fmt.Println("Welcome to Tic-Tac-Toe")
}

func (g *Game) GameLogic() {
	g.board.PrintBoard()

	player := ChoosePlayer(g.numberOfTurns)

	err := g.board.PlaceOnBoard(player, ChooseLocation(player))
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	} else {
		g.IncrementGameTurns()
	}

	g.gameOver = g.board.CheckBoardForWinner()
}

func (g *Game) IncrementGameTurns() {
	g.numberOfTurns++
}
