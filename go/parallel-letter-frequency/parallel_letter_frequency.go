package letter

func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap, len(texts))
	for _, text := range texts {
		go func(text string) {
			c <- Frequency(text)
		}(text)
	}

	// Combine results from separate goroutines
	result := FreqMap{}
	for i := 0; i < len(texts); i++ {
		for k, v := range <-c {
			result[k] += v
		}
	}

	return result
}
