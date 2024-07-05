package service

import "fmt"

type Game struct {
	board         *Board
	numberOfTurns int
	gameOver      bool
	currentPlayer string
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

	g.GameOverPrompt(g.currentPlayer)
}

func WelcomePrompt() {
	fmt.Println("Welcome to Tic-Tac-Toe")
}

func (g *Game) GameLogic() {
	g.board.PrintBoard()

	g.currentPlayer = ChoosePlayer(g.numberOfTurns)

	var loc int
	fmt.Print("enter location: ")
	fmt.Scanln(&loc)

	err := g.board.PlaceOnBoard(g.currentPlayer, ChooseLocation(g.currentPlayer, loc))
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

func (g *Game) GameOverPrompt(player string) {
	fmt.Printf("Game Over, player %s won!\n", player)
	fmt.Printf("Number of turns: %d\n\n", g.numberOfTurns)

	g.board.PrintBoard()
}
