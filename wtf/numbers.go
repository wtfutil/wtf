package wtf

import "math"

// Round rounds a float to an integer
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// TruncateFloat64 truncates the decimal places of a float64 to the specified precision
func TruncateFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
