package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate will take a string and convert it to an acronym
func Abbreviate(s string) string {
	acronym := ""
	useNext := false
	var lastRune rune
	for i, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			// if we find a word break, then use the next character
			// for the acronym
			useNext = true
		} else if useNext {
			acronym += strings.ToUpper(s[i : i+1])
			useNext = false
		} else if unicode.IsUpper(r) && !unicode.IsUpper(lastRune) {
			// If we find an upper case character, consider that to be
			// part of the acronym unless it's all caps
			acronym += s[i : i+1]
		}
		lastRune = r
	}

	return acronym
}
