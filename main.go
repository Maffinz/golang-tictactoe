package main

import "maffniz.com/tictactoe/service"

func main() {
	myGame := service.InitializeGame()

	myGame.GameLoop()

}
