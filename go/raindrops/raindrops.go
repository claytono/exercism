package raindrops

import "strconv"

const testVersion = 2

// Convert will print a silly message depending on the factors of the input
func Convert(i int) string {
	output := ""
	if i%3 == 0 {
		output += "Pling"
	}

	if i%5 == 0 {
		output += "Plang"
	}

	if i%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		return strconv.Itoa(i)
	}

	return output
}
