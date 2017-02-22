package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func trimNonLetters(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, s)
}

func Hey(question string) string {
	question = strings.TrimSpace(question)

	qlen := len(question)
	if qlen == 0 {
		return "Fine. Be that way!"
	}

	qLetters := trimNonLetters(question)
	if len(qLetters) > 0 {
		if qLetters == strings.ToUpper(qLetters) {
			return "Whoa, chill out!"
		}
	}

	if qlen > 0 && question[qlen-1] == '?' {
		return "Sure."
	}

	return "Whatever."
}
