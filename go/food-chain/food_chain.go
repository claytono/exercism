package foodchain

import (
	"strings"
)

const testVersion = 3

type verseInfoStruct struct {
	// The animal that was swallowed
	animal string
	// Extra sentence to print immediately after announcing the swallowing
	swallowedSentence string
	// Extra text to add after the animal name when saying what she was trying
	// to catch after the initial announcement.
	flavorText string
	// If true, then the swallowing killed the old lady, exit early.
	diedEarly bool
}

var verseInfo = []verseInfoStruct{
	{"fly", "", "", false},
	{"spider", "It wriggled and jiggled and tickled inside her.",
		" that wriggled and jiggled and tickled inside her", false},
	{"bird", "How absurd to swallow a bird!", "", false},
	{"cat", "Imagine that, to swallow a cat!", "", false},
	{"dog", "What a hog, to swallow a dog!", "", false},
	{"goat", "Just opened her throat and swallowed a goat!", "", false},
	{"cow", "I don't know how she swallowed a cow!", "", false},
	{"horse", "She's dead, of course!", "", true},
}

func Verse(n int) string {
	if n > len(verseInfo) {
		return ""
	}

	offset := n - 1
	lines := []string{
		"I know an old lady who swallowed a " + verseInfo[offset].animal + ".",
	}
	if len(verseInfo[offset].swallowedSentence) > 0 {
		lines = append(lines, verseInfo[offset].swallowedSentence)
	}

	if verseInfo[offset].diedEarly {
		return strings.Join(lines, "\n")
	}

	for i := offset; i > 0; i-- {
		line := "She swallowed the " + verseInfo[i].animal + " to catch the " +
			verseInfo[i-1].animal + verseInfo[i-1].flavorText + "."
		lines = append(lines, line)
	}

	lines = append(lines, "I don't know why she swallowed the fly. Perhaps she'll die.")

	return strings.Join(lines, "\n")
}

func Verses(to, from int) string {
	verses := make([]string, 0, from-to+1)

	for i := to; i <= from; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, len(verseInfo))
}
