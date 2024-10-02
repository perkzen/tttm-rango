package game

import (
	"fmt"
	"github.com/perkzen/tttm-go/pkg/utils"
	"strconv"
	"strings"
)

type Symbol string

const (
	X     Symbol = "X"
	O     Symbol = "O"
	EMPTY Symbol = ""
)

func ToPlayerSymbol(s string) (Symbol, error) {
	switch s {
	case "X":
		return X, nil
	case "O":
		return O, nil
	default:
		return EMPTY, fmt.Errorf("invalid player symbol: %s", s)
	}
}

type Board [][]Symbol

type Game struct {
	Board Board
	size  int
}

func NewGame(size int, moves string) *Game {
	board := movesToBoard(size, moves)
	return &Game{board, size}
}

func movesToBoard(size int, moves string) Board {
	board := newEmptyBoard(size)

	moveList := strings.Split(moves, "-")

	for _, move := range moveList {
		parts := strings.Split(move, "_")

		if len(parts) != 3 {
			continue
		}

		symbol, err := ToPlayerSymbol(parts[0])
		if err != nil {
			continue
		}

		row, rowErr := strconv.Atoi(parts[1])
		col, colErr := strconv.Atoi(parts[2])

		if rowErr != nil || colErr != nil || row < 0 || row >= size || col < 0 || col >= size {
			continue
		}

		board[row][col] = symbol
	}

	return board
}

func newEmptyBoard(size int) Board {
	board := make(Board, size)

	for i := range board {
		board[i] = make([]Symbol, size)
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}

	return board
}

func (g *Game) GetRandomMove(player Symbol) (int, int) {
	for {
		row := utils.RandomInt(0, g.size)
		col := utils.RandomInt(0, g.size)

		if row >= 0 && row < g.size && col >= 0 && col < g.size && g.Board[row][col] == EMPTY {
			g.Board[row][col] = player
			return row, col
		}
	}
}
