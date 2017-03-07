package lsproduct

import (
	"errors"
	"fmt"
	"unicode"
)

const testVersion = 4

func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	runes := []rune(digits)

	if span < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	if len(runes) < span {
		return 0, errors.New("span cannot be longer than the number")
	}

	largest := int64(0)
	for i := 0; i <= len(runes)-span; i++ {
		product := int64(1)
		for j := i; j < i+span; j++ {
			if !unicode.IsNumber(runes[j]) {
				return 0, fmt.Errorf("string '%v' contains non-digits", digits)
			}
			// Converting the digit ourselves here is nearly 10x faster than
			// calling Atoi.  Normally I'd just call Atoi here, but it doesn't
			// handle unicode numbers properly either, and this exercise has a
			// benchmark test!
			product *= int64(runes[j] - '0')
		}

		if product > largest {
			largest = product
		}
	}

	return largest, nil
}
