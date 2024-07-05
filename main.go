package main

import "github.com/Maffinz/tictactoe/service"

func main() {
	myGame := service.InitializeGame()

	myGame.GameLoop()

}
