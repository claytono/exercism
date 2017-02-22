package hamming

import (
	"errors"
)

const testVersion = 5

// Distance will calculated the Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := errors.New("Strings must be of the same length to calculate distance")
		return 0, err
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
