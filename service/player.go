package service

import (
	"github.com/Maffinz/golang-tictactoe/model"
)

func ChoosePlayer(t int) string {
	if t%2 == 0 {
		return "O"
	}

	return "X"
}

func ChooseLocation(player string, loc int) model.Cordinate {
	return model.Cordinate{X: calculateX(loc), Y: calculateY(loc)}
}

func calculateY(loc int) int {
	// 3 Has to to change to a variable to accomodate for other grids
	if loc < 3 {
		return loc - 1
	}

	if loc%3 == 0 {
		return 2
	}
	return loc - (3 * (loc / 3)) - 1
}
func calculateX(loc int) int {
	// 3 Has to change to a variable to accomodate for other grids
	if loc < 3 {
		return 0
	}

	if loc%3 != 0 {
		return loc / 3
	}
	return (loc / 3) - 1
}
