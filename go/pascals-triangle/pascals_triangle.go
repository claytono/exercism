package pascal

// Given a previous row, calculate the next row
func nextRow(row []int) []int {
	next := []int{row[0]}

	for i, v := range row {
		if i != 0 {
			next = append(next, row[i-1]+v)
		}
	}

	return append(next, row[len(row)-1])
}

// Triangle will calculate all levels of a pascal's triangle of n depth
func Triangle(n int) [][]int {
	result := [][]int{{1}}

	for i := 1; i < n; i++ {
		result = append(result, nextRow(result[i-1]))
	}

	return result
}
