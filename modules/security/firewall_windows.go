//go:build windows

package security

// firewallStateLinux is a no-op stub on Windows since Linux firewall checks
// use syscall.WaitStatus which is not available on Windows.
func firewallStateLinux() string {
	return ""
}

func firewallStealthStateLinux() string {
	return "[white]N/A[white]"
}
