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
	b := &Board{W: w, H: h, alive: true}
	b.grid = make([][]bool, h)
	for y := 0; y < h; y++ {
		b.grid[y] = make([]bool, w)
		for x := 0; x < w; x++ {
			b.grid[y][x] = true
		}
	}
	return b
}

func (b *Board) IsUneaten(x, y int) bool {
	return inBounds(b, x, y) && b.grid[y][x]
}

func (b *Board) Chomp(x, y int) (atePoison bool, err error) {
	if !inBounds(b, x, y) {
		return false, errors.New("move out of bounds")
	}
	if !b.grid[y][x] {
		return false, errors.New("that square is already eaten")
	}
	for j := y; j < b.H; j++ {
		for i := x; i < b.W; i++ {
			b.grid[j][i] = false
		}
	}
	if !b.grid[0][0] { // poison eaten :P
		b.alive = false
		return true, nil
	}
	return false, nil
}

func (b *Board) HasAnyMove() bool {
	for y := 0; y < b.H; y++ {
		for x := 0; x < b.W; x++ {
			if b.grid[y][x] {
				return true
			}
		}
	}
	return false
}

func (b *Board) Draw() {
	fmt.Print("    ")
	for x := 0; x < b.W; x++ {
		fmt.Printf("%2d ", x+1)
	}
	fmt.Println()
	fmt.Print("    ")
	for x := 0; x < b.W; x++ {
		fmt.Print("───")
	}
	fmt.Println()
	for y := 0; y < b.H; y++ {
		fmt.Printf("%2d │ ", y+1)
		for x := 0; x < b.W; x++ {
			switch {
			case x == 0 && y == 0 && b.grid[y][x]:
				fmt.Printf("%c  ", cellPoison)
			case b.grid[y][x]:
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
