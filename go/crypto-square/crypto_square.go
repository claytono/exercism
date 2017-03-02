package cryptosquare

import (
	"math"
	"unicode"
)

const testVersion = 2

// normalize modifies the plain-text string given remove all runes other than
// letters and numbers and convert all letters to lower case.  This is for use
// by the Encode func.
func normalize(pt string) []rune {
	output := make([]rune, 0, len(pt))
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			output = append(output, unicode.ToLower(r))
		}
	}

	return output
}

// calcSquareDimensions takes a length of a string, and returns the number of
// rows and columns that would are needed to wrap it into a square.  If a
// perfect square cannot be found, the columns in the rectangle will be no
// more than one more than the number of rows.
func calcSquareDimensions(length int) (r, c int) {
	// Get the square root of the string length, and round up to the nearest
	// integer.  This will give the number of rows we need.
	r = int(math.Sqrt(float64(length)) + 0.5)

	// The number of columns will just be the length divided by the number of rows
	c = length / r

	// If the number of rows and columns isn't exactly the length, then we
	// have some extra characters at the end, so add one more column
	if c*r != length {
		c++
	}

	return
}

func Encode(pt string) string {
	normal := normalize(pt)
	normalLen := len(normal)

	if normalLen == 0 {
		return ""
	}

	rowCount, colCount := calcSquareDimensions(normalLen)

	ct := make([]rune, 0, len(normal))

	// Iterate over the rows and columns.  We may have an incomplete column at
	// the end, so validate we don't run off the end of the plaintext string
	for c := 0; c < colCount; c++ {
		for r := 0; r < rowCount; r++ {
			origPos := r*colCount + c
			if origPos >= len(normal) {
				break
			}
			ct = append(ct, normal[origPos])
		}
		ct = append(ct, ' ')
	}

	// Remove extra trailing space
	ct = ct[:len(ct)-1]

	return string(ct)
}
