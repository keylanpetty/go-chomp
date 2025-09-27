package chomp

import (
	"fmt"
)

func Welcome(board *Board) {
	fmt.Println("Welcome to my chomp game! Invented by David Gale, Chomp is a two-player, mathematical strategy game played on a rectangular grid,\nor a ~chocolate bar~, where the player forced to eat the poisoned top-left square loses")
	fmt.Printf("\nCHOMP — %dx%d (poison chocolate block top-left)\n", board.W, board.H)
	fmt.Println("Pick a square to chomp: that square and everything bottom-right disappears.")
	fmt.Println("Whoever eats poisonous chocolate leaves their opponent victorious. Coordinates are COLUMN ROW (1-based), e.g., `3 2` or `c2`.")
	fmt.Println("Type `help`, `board`, or `quit`")
	fmt.Println()
}
