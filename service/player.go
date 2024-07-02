package service

import (
	"fmt"
	"tictactoe/TicTacToe/model"
)

func ChoosePlayer(t int) string {
	if t%2 == 0 {
		return "O"
	}

	return "X"
}

func ChooseLocation(player string) model.Cordinate {
	c := model.Cordinate{}

	fmt.Printf("Player '%s' enter Cordinate [x,y]: ", player)
	fmt.Scanln(&c.X, &c.Y)

	return c
}
