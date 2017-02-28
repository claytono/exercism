package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				sum += i
				// Each number under the limit should only be added to the sum
				// once, even if it's divisible by multiple divisors
				break
			}
		}
	}

	return
}
