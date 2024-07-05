package main

import "github.com/Maffinz/golang-tictactoe/service"

func main() {
	myGame := service.InitializeGame()

	myGame.GameLoop()

}
