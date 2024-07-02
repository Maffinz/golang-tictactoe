package main

import "tictactoe/TicTacToe/service"

func main() {
	myGame := service.InitializeGame()

	myGame.GameLoop()
}
