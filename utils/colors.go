package utils

import "fmt"

// ColorizePercent provides a standard way to colorize percentages for which
// large numbers are good (green) and small numbers are bad (red).
func ColorizePercent(percent float64) string {
	var color string

	switch {
	case percent >= 70:
		color = "green"
	case percent >= 35:
		color = "yellow"
	case percent < 0:
		color = "grey"
	default:
		color = "red"
	}

	return fmt.Sprintf("[%s]%v[%s]", color, percent, "white")
}
