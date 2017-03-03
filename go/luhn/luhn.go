package luhn

import (
	"unicode"
)

const testVersion = 1

// Valid calculates a Luhn checksum across the digits in the string and
// returns true if the digits checksum is valid.
func Valid(n string) bool {
	runes := []rune(n)
	digits := make([]int, 0, len(n))
	// Walk string backwards, so we can easily double every second digit from
	// the end when we're done collecting digits
	for i := len(n) - 1; i >= 0; i-- {
		if unicode.IsSpace(runes[i]) {
			continue
		}
		if !unicode.IsDigit(runes[i]) {
			return false
		}

		// Convert the digit to a integer and put it on the end.
		digits = append(digits, int(runes[i]-'0'))
	}

	if len(digits) <= 1 {
		return false
	}

	// Since the int array is now in reverse order, traverse the entire array,
	// adding even entries and doubling then adding odd ones.
	sum := 0
	for i, x := range digits {
		if i%2 == 0 {
			sum += x
		} else if x*2 > 9 {
			sum += x*2 - 9
		} else {
			sum += x * 2
		}
	}

	return sum%10 == 0
}
