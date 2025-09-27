package chomp

import "math/rand"

type move struct {
	x     int
	y     int
	score int
}

func CPUChoose(rng *rand.Rand, board *Board, difficulty bool) (int, int) {

	if onlyPoisonLeft(board) {
		return 0, 0
	}

	var moves []move

	// Difficulty: EASY
	// algo for a simple greedy bite - CPU will see the best move as the largest chomp in leftmost column
	for y := 0; y < board.H; y++ {
		for x := 0; x < board.W; x++ {
			if !board.grid[y][x] {
				continue
			}
			if x == 0 && y == 0 {
				continue
			}
			area := (board.W - x) * (board.H - y)
			penalty := 1
			if 1-x > 0 {
				penalty += 1 - x
			}
			if 1-y > 0 {
				penalty += 1 - y
			}
			moves = append(moves, move{x, y, area - penalty})
		}
	}

	// TODO: Difficulty: HARD
	// recursive DFS of winning moves - think rust project....


	// TODO: shuffle can panic -> make a defer func to see recovery
	rng.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	best := moves[0]
	for _, m := range moves {
		if m.score > best.score {
			best = m
		}
	}
	return best.x, best.y
}

func onlyPoisonLeft(board *Board) bool {
	for y := 0; y < board.H; y++ {
		for x := 0; x < board.W; x++ {
			if board.grid[y][x] && !(x == 0 && y == 0) {
				return false
			}
		}
	}
	return board.grid[0][0]
}
