package chomp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrQuit = errors.New("quit")

// ParseMove parses "3 2" or "b4" into 0-based (x,y) validated for the board
func ParseMove(input string, b *Board) (int, int, error) {
	s := strings.TrimSpace(input)
	switch strings.ToLower(s) {
	case "q", "quit", "exit":
		return 0, 0, ErrQuit
	case "h", "help", "?":
		return 0, 0, fmt.Errorf(
			"choose COLUMN ROW (1-based), e.g., `3 2` or `c2`. Poison is at (1,1)",
		)
	case "b", "board":
		return 0, 0, fmt.Errorf("board")
	}

	if x, y, ok := tryNumeric(s); ok {
		x--
		y--
		if !inBounds(b, x, y) {
			return 0, 0, fmt.Errorf("out of bounds")
		}
		if !b.IsUneaten(x, y) {
			return 0, 0, fmt.Errorf("already eaten")
		}
		return x, y, nil
	}
	if x, y, ok := tryAlphaNum(s); ok {
		x--
		y--
		if !inBounds(b, x, y) {
			return 0, 0, fmt.Errorf("out of bounds")
		}
		if !b.IsUneaten(x, y) {
			return 0, 0, fmt.Errorf("already eaten")
		}
		return x, y, nil
	}
	return 0, 0, fmt.Errorf("couldn't parse; try `3 2` or `b4`")
}

func tryNumeric(s string) (int, int, bool) {
	parts := strings.Fields(s)
	if len(parts) != 2 {
		return 0, 0, false
	}
	a, err1 := strconv.Atoi(parts[0])
	b, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0, false
	}
	return a, b, true
}

func tryAlphaNum(s string) (int, int, bool) {
	s = strings.ToLower(strings.TrimSpace(s))
	if len(s) < 2 {
		return 0, 0, false
	}
	i := 0
	for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
		i++
	}
	if i == 0 || i == len(s) {
		return 0, 0, false
	}
	colStr := s[:i]
	rowStr := s[i:]
	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return 0, 0, false
	}
	col := 0
	for _, r := range colStr {
		if r < 'a' || r > 'z' {
			return 0, 0, false
		}
		col = col*26 + int(r-'a'+1)
	}
	return col, row, true
}
