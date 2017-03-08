package sieve

const prime = false
const composite = true

// Sieve implements the Sieve of Eratosthenes
func Sieve(limit int) []int {
	c := make([]bool, limit)

	// We only have to test numbers up to half the limit because all of them
	// will be multipied by at least 2
	for p := 2; p < limit/2; p++ {
		if c[p] == composite {
			continue
		}

		for m := 2; m*p < limit; m++ {
			if c[p*m] == prime {
				c[p*m] = composite
			}
		}
	}

	primes := []int{}
	for i := 2; i < limit; i++ {
		if c[i] == prime {
			primes = append(primes, i)
		}
	}

	return primes
}
