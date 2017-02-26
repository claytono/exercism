package etl

import "strings"

const testVersion = 1

func Transform(old map[int][]string) (result map[string]int) {
	result = make(map[string]int, 26)
	for score, letters := range old {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return
}
