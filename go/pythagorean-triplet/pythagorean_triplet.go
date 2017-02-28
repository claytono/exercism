package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) []Triplet {
	results := []Triplet{}

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			// Calculate what c would be
			cf := math.Sqrt(float64(a*a + b*b))

			// Convert back to an integer, and then compare to the float
			// version.  If they're not equal then the length isn't an integer
			// and isn't valid a valid solution
			c := int(cf)
			if c <= max && float64(c) == cf {
				results = append(results, Triplet{a, b, c})
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
