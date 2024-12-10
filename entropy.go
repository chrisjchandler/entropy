package entropy

import (
	"math"
)

// Calculate computes the Shannon entropy of a given input string.
func Calculate(input string) float64 {
	// Count the frequency of each character
	frequency := make(map[rune]float64)
	for _, char := range input {
		frequency[char]++
	}

	// Calculate the total number of characters
	total := float64(len(input))

	// Calculate entropy
	var entropy float64
	for _, count := range frequency {
		probability := count / total
		entropy += probability * math.Log2(probability)
	}

	return -entropy // Negate the sum as entropy is positive
}
