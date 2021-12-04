//go:build windows

package system

import (
	"os/exec"
	"strings"
)

type SystemInfo struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

func NewSystemInfo() *SystemInfo {
	m := make(map[string]string)

	cmd := exec.Command("powershell.exe", "(Get-CimInstance Win32_OperatingSystem).version")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	s := strings.Split(string(out), ".")
	m["ProductName"] = "Windows"
	m["ProductVersion"] = "Windows " + s[0] + "." + s[1]
	m["BuildVersion"] = s[2]

	sysInfo := SystemInfo{
		ProductName:    m["ProductName"],
		ProductVersion: m["ProductVersion"],
		BuildVersion:   m["BuildVersion"],
	}

	return &sysInfo
}
