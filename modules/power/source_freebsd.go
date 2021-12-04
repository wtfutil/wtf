//go:build freebsd

package power

// powerSource returns the name of the current power source, probably one of
// "AC Power" or "Battery Power"
func powerSource() string {
	switch batteryState {
	case "1":
		return "AC Power"
	case "0":
		return "Battery Power"
	}
	return batteryState
}
