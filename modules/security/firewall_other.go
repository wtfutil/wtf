//go:build !darwin && !linux && !windows

package security

func FirewallState() string {
	return ""
}

func FirewallStealthState() string {
	return ""
}
