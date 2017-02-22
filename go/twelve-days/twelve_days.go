package twelve

import "fmt"

const testVersion = 1

// Maps days of Christmas to the gift added on that day
var gifts = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

// Maps numbers to the original word representing them
var numberToOrdinal = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func getNextGift(day int) string {
	if day == 1 {
		return gifts[1]
	}

	if day == 2 {
		// Have to special case day two in order to insert the and conjunction
		// between the second and first days.  Can't do this on day one, since
		// it might be the only day called.
		return gifts[day] + ", and " + getNextGift(1)
	}

	return gifts[day] + ", " + getNextGift(day-1)

}

// Verse returns a single verse of the song "The Twelve Days of Christmas" when
// given the day of the verse.
func Verse(day int) string {
	leadIn := fmt.Sprintf("On the %s day of Christmas my true love gave to me, ", numberToOrdinal[day])
	return leadIn + getNextGift(day) + "."
}

// Song returns the complete lyrics of the song "The Twelve Days  of Christmas"
func Song() string {
	song := ""
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}

	return song
}
