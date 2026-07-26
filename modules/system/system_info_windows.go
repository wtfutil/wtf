//go:build windows

package system

import (
	"os/exec"
	"strconv"
	"strings"
)

type SystemInfo struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

// win11BuildNumber is the first Windows build number reported as Windows 11.
// Windows 11 kept the "10.0" major.minor version, so the build number is the
// only reliable signal Win32_OperatingSystem gives us to tell it apart from
// Windows 10.
const win11BuildNumber = 22000

func NewSystemInfo() *SystemInfo {
	cmd := exec.Command("powershell.exe", "(Get-CimInstance Win32_OperatingSystem).version")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	productVersion, buildVersion := windowsProductVersion(string(out))

	return &SystemInfo{
		ProductName:    "Windows",
		ProductVersion: productVersion,
		BuildVersion:   buildVersion,
	}
}

// windowsProductVersion turns the raw "major.minor.build" string reported by
// Win32_OperatingSystem into a human-friendly product version and separate
// build number. Windows 11 kept the "10.0" major.minor version, so the build
// number is the only reliable signal available to distinguish it from
// Windows 10.
func windowsProductVersion(rawVersion string) (productVersion, build string) {
	s := strings.Split(strings.TrimSpace(rawVersion), ".")
	if len(s) < 3 {
		return "Windows " + strings.TrimSpace(rawVersion), ""
	}

	major, minor, build := s[0], s[1], s[2]

	if major == "10" {
		if buildNum, err := strconv.Atoi(build); err == nil && buildNum >= win11BuildNumber {
			return "Windows 11", build
		}
		return "Windows 10", build
	}

	return "Windows " + major + "." + minor, build
}
