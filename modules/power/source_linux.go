//go:build linux

package power

// powerSource returns the name of the current power source, probably one of
// "AC Power" or "Battery Power"
func powerSource() string {
	switch batteryState {
	case "charging", "fully-charged":
		return "AC Power"
	case "discharging":
		return "Battery Power"
	}
	return batteryState
}
