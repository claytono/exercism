package house

import "strings"

const testVersion = 1

type verse struct {
	verb string
	noun string
}

var verseParts = []verse{
	{"lay in", "house that Jack built"},
	{"ate", "malt"},
	{"killed", "rat"},
	{"worried", "cat"},
	{"tossed", "dog"},
	{"milked", "cow with the crumpled horn"},
	{"kissed", "maiden all forlorn"},
	{"married", "man all tattered and torn"},
	{"woke", "priest all shaven and shorn"},
	{"kept", "rooster that crowed in the morn"},
	{"belonged to", "farmer sowing his corn"},
	{"", "horse and the hound and the horn"},
}

// getVerseLine returns the specific line in a verse.  n is zero based.
func getVerseLine(n int, first bool) string {
	v := verseParts[n]
	if first {
		return "This is the " + v.noun
	}
	return "that " + v.verb + " the " + v.noun
}

func Verse(n int) string {
	lines := []string{}
	for i := n; i > 0; i-- {
		first := (i == n)
		lines = append(lines, getVerseLine(i-1, first))
	}
	return strings.Join(lines, "\n") + "."
}

func Song() string {
	verses := []string{}
	for i := 0; i < 12; i++ {
		verses = append(verses, Verse(i+1))
	}
	return strings.Join(verses, "\n\n")
}
