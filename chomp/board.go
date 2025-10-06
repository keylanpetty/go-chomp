package chomp

import (
	"errors"
	"fmt"
)

const (
	cellUneaten = '🍫'
	cellGone    = '⬛'
	cellPoison  = '☠'
)

type Board struct {
	W, H  int
	grid  [][]bool
	alive bool
}

func (b *Board) PoisonEaten() bool {
	return !b.alive
}

func NewBoard(w, h int) *Board {
	board := &Board{W: w, H: h, alive: true}
	board.grid = make([][]bool, h)
	for y := 0; y < h; y++ {
		board.grid[y] = make([]bool, w)
		for x := 0; x < w; x++ {
			board.grid[y][x] = true
		}
	}
	return board
}

func (board *Board) IsUneaten(x, y int) bool {
	return inBounds(board, x, y) && board.grid[y][x]
}

func (board *Board) Chomp(x, y int) (atePoison bool, err error) {
	if !inBounds(board, x, y) {
		return false, errors.New("move out of bounds")
	}
	if !board.grid[y][x] {
		return false, errors.New("that square is already eaten")
	}
	for j := y; j < board.H; j++ {
		for i := x; i < board.W; i++ {
			board.grid[j][i] = false
		}
	}
	if !board.grid[0][0] { // poison eaten :P
		board.alive = false
		return true, nil
	}
	return false, nil
}

func (board *Board) HasAnyMove() bool {
	for y := 0; y < board.H; y++ {
		for x := 0; x < board.W; x++ {
			if board.grid[y][x] {
				return true
			}
		}
	}
	return false
}

func (board *Board) Draw() {
	fmt.Print("    ")
	for x := 0; x < board.W; x++ {
		fmt.Printf("%2d ", x+1)
	}
	fmt.Println()
	fmt.Print("    ")
	for x := 0; x < board.W; x++ {
		fmt.Print("───")
	}
	fmt.Println()
	for y := 0; y < board.H; y++ {
		fmt.Printf("%2d │ ", y+1)
		for x := 0; x < board.W; x++ {
			switch {
			case x == 0 && y == 0 && board.grid[y][x]:
				fmt.Printf("%c  ", cellPoison)
			case board.grid[y][x]:
				fmt.Printf("%c  ", cellUneaten)
			default:
				fmt.Printf("%c  ", cellGone)
			}
		}
		fmt.Println()
	}
}

func inBounds(board *Board, x, y int) bool {
	return x >= 0 && x < board.W && y >= 0 && y < board.H
}
