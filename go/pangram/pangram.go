package pangram

import (
	"unicode"
)

const testVersion = 1

func isASCIILetter(r rune) bool {
	return unicode.IsLetter(r) && r < unicode.MaxASCII
}

// IsPangram returns true of the string provided contains all ASCII letters,
// but ignores all other characters.
func IsPangram(s string) bool {
	letters := map[rune]bool{}

	for _, r := range s {
		if !isASCIILetter(r) {
			continue
		}
		letters[unicode.ToUpper(r)] = true
	}
	return len(letters) == 26
}
