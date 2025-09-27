package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/keylanpetty/go-chomp/chomp"
)

type playerKind int

const (
	human playerKind = iota
	cpu
)

type player struct {
	kind playerKind
	name string
}

func main() {
	// Local range as per stdlib  (no global rand.Seed??)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	w := flag.Int("w", 6, "width")
	h := flag.Int("h", 5, "height")
	vsCPU := flag.Bool("cpu", false, "play against the computer")
	flag.Parse()

	if *w < 2 || *h < 2 {
		fmt.Println("Minimum size is 2x2 to make it interesting... and tasty...")
		return
	}

	board := chomp.NewBoard(*w, *h)

	p1 := player{human, "Player 1"}
	var p2 player
	if *vsCPU {
		p2 = player{cpu, "CPU"}
	} else {
		p2 = player{human, "Player 2"}
	}

	chomp.Welcome(board)
	board.Draw()

	sc := bufio.NewScanner(os.Stdin)
	turn := 0

	for {
		// Whose turn is it anyway?
		pl := p1
		opp := p2
		if turn%2 == 1 {
			pl = p2
			opp = p1
		}

		fmt.Printf("\n%s, your move → ", pl.name)

		moved := false
		var x, y int

		if pl.kind == cpu {
			time.Sleep(500 * time.Millisecond) // a little drama
			x, y = chomp.CPUChoose(rng, board, false)
			fmt.Printf("%d %d\n", x+1, y+1)
			moved = true
		} else {
			for {
				if !sc.Scan() { // EOF or closed stdin
					fmt.Println("Bye!")
					return
				}
				xx, yy, err := chomp.ParseMove(sc.Text(), board)
				if err != nil {
					switch {
					case errors.Is(err, chomp.ErrQuit):
						fmt.Println("The poison chocolate block is at large... but first! Errors!")
						return
					case err.Error() == "board":
						board.Draw()
						fmt.Print("the board aint changin but i like your sense of whimsy... now quickly! chomp! ")
						continue
					default:
						fmt.Println(" ", err)
						fmt.Print(" your move cowboy... ")
						continue
					}
				}
				x, y = xx, yy
				moved = true
				break
			}
		}

		// Only apply a move if we actually got one.
		if !moved {
			continue
		}

		if atePoison, err := board.Chomp(x, y); err != nil {
			fmt.Println(" ", err)
			continue
		} else if atePoison {
			// End the game immediately with the correct winner/loser !!!

			fmt.Println("Uh oh... that didn't go down well...")
			fmt.Printf("\n%s ate the poison choco-block. Welcome to ghost-hood my friend. \n\n%s continues to risk their life in the name of sweet treats! A noble pursuit!\n", pl.name, opp.name)
			fmt.Println("Good game!")
			return
		}

		board.Draw()
		turn++
	}
}
