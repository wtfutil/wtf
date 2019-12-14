package utils

// SumInts takes a slice of ints and returns the sum of them
func SumInts(vals []int) int {
	sum := 0

	for _, a := range vals {
		sum += a
	}

	return sum
}
