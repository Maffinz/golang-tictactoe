package service

import (
	"errors"
	"fmt"

	"github.com/Maffinz/tictactoe/model"
)

type Board struct {
	board *[][]string
}

func initializeBoard() *Board {
	return &Board{
		board: &[][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
	}
}

func BoardWinVertical(b *[][]string) bool {
	var won bool
	for i := 0; i < len((*b)[0]); i++ {
		won = true
		charFound := ""
		for j := 0; j < len(*b); j++ {
			charFound = (*b)[0][i]

			if (*b)[j][i] != charFound {
				won = false
			}
		}

		if won {
			return true
		}
	}

	return won
}

func BoardWinHorizontal(b *[][]string) bool {
	won := true

	for _, row := range *b {
		charFound := row[0]
		won = true
		for _, val := range row {
			if val != charFound {
				won = false
			}
		}

		if won {
			return true
		}
	}

	return won
}

func BoardWinDiagonal(b *[][]string) bool {
	won := true
	charFound := (*b)[0][0]

	// Right to Left
	for i := 0; i < len(*b); i++ {
		if (*b)[i][i] != charFound {
			won = false
		}
	}

	if won {
		return won
	}

	won = true
	charFound = (*b)[0][2]
	// Left to Right
	for i := 2; i >= 0; i-- {
		if (*b)[2-i][i] != charFound {
			won = false
		}
	}

	return won
}

func (b *Board) CheckBoardForWinner() bool {
	return BoardWinVertical(b.board) || BoardWinHorizontal(b.board) || BoardWinDiagonal(b.board)
}

func (b *Board) PlaceOnBoard(player string, cor model.Cordinate) error {
	if cor.X >= len(*b.board) {
		return errors.New("x is greater than bounds")
	}
	if cor.Y >= len((*b.board)[0]) {
		return errors.New("y is greater than bounds")
	}

	if (*b.board)[cor.X][cor.Y] == "X" || (*b.board)[cor.X][cor.Y] == "O" {
		return errors.New("player already in position")
	}

	(*b.board)[cor.X][cor.Y] = player
	return nil
}

func (b *Board) PrintBoard() {
	for _, row := range *(b.board) {
		for _, val := range row {
			fmt.Print(val, "|")
		}
		fmt.Println()
		fmt.Println("------")
	}
}
