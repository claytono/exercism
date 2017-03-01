package isogram

import "unicode"

const testVersion = 1

func IsIsogram(word string) bool {
	// Due to unicode support, this map may grow over 26 entries, but this is
	// a reasonable starting point.
	lettersFound := make(map[rune]bool, 26)

	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if lettersFound[r] {
			return false
		}

		lettersFound[r] = true
	}

	return true
}
