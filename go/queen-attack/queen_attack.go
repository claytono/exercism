package queenattack

import (
	"errors"
	"math"
)

const testVersion = 2

func parsePosition(pos string) (x, y int, err error) {
	if len(pos) != 2 {
		return 0, 0, errors.New("position must be 2 chacters long")
	}

	if pos[0] < 'a' || pos[0] > 'h' {
		return 0, 0, errors.New("position must start with a through h")
	}
	x = int(pos[0]) - 'a'

	if pos[1] < '1' || pos[1] > '8' {
		return 0, 0, errors.New("position must end with 1 through 8")
	}
	y = int(pos[1]) - '1'

	return x, y, nil
}

// CanQueenAttack takes the algebraic chess notation of the white queen and
// the black piece and returns whether or not the queen is in a valid attack
// position.
func CanQueenAttack(w, b string) (bool, error) {
	var wx, wy, bx, by int
	var err error

	if wx, wy, err = parsePosition(w); err != nil {
		return false, err
	}

	if bx, by, err = parsePosition(b); err != nil {
		return false, err
	}

	if wx == bx && wy == by {
		return false, errors.New("black and white cannot be in the same location")
	}

	// Check diagonal moves
	if math.Abs(float64(wx-bx)) == math.Abs(float64(wy-by)) {
		return true, nil
	}

	// Check horizonal and vertical moves
	if wx-bx == 0 || wy-by == 0 {
		return true, nil
	}

	return false, nil
}
