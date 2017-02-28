package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	results := []Triplet{}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			// My original solution used float math to calculate exactly what
			// c would be, then checked if it was under max.  That solution is
			// probably faster for large enough values of max, but the test
			// case given only goes up to 20.  In that case, the brute force
			// approach is faster.
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					results = append(results, Triplet{a, b, c})
				}
			}
		}
	}
	return results
}

func Sum(p int) []Triplet {
	results := []Triplet{}
	candidates := Range(1, p)
	for _, c := range candidates {
		if c[0]+c[1]+c[2] == p {
			results = append(results, c)
		}
	}

	return results
}
