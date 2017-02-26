package grains

import "errors"

const testVersion = 1

// Square returns the total number of grains on the specific square number
// provided.
func Square(x int) (result uint64, err error) {
	if x <= 0 || x > 64 {
		err = errors.New("square number must be between 1 and 64")
		return
	}

	result = 1 << uint(x-1)
	return
}

// Total sums the grains on the entire board and returns the total.
func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		x, _ := Square(i)
		total += x
	}
	return
}
