package security

import (
	"os/exec"
	"runtime"

	"github.com/senorprogrammer/wtf/wtf"
)

// https://github.com/yelinaung/wifi-name/blob/master/wifi-name.go
const osxWifiCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxWifiArg = "-I"

/* -------------------- Exported Functions -------------------- */

func wifiEncryptionLinux() string {
	cmd := exec.Command("nmcli", "-t", "-f", "active,security", "dev", "wifi")
	out := wtf.ExecuteCommand(cmd)
	name := wtf.FindMatch(`yes:(.+)`, out)
	if len(name) > 0 {
		return name[0][1]
	}
	return ""
}

func wifkEncryptionMacOS() string {
	name := wtf.FindMatch(`s*auth: (.+)s*`, wifiInfo())
	return matchStr(name)
}

func WifiEncryption() string {
	switch runtime.GOOS {
	case "linux":
		return wifiEncryptionLinux()
	case "macos":
		return wifkEncryptionMacOS()
	default:
		return ""
	}
}

func wifiNameMacOS() string {
	name := wtf.FindMatch(`s*SSID: (.+)s*`, wifiInfo())
	return matchStr(name)
}

func wifiNameLinux() string {
	cmd := exec.Command("nmcli", "-t", "-f", "active,ssid", "dev", "wifi")
	out := wtf.ExecuteCommand(cmd)
	name := wtf.FindMatch(`yes:(.+)`, out)
	if len(name) > 0 {
		return name[0][1]
	}
	return ""
}

func WifiName() string {
	switch runtime.GOOS {
	case "linux":
		return wifiNameLinux()
	case "macos":
		return wifiNameMacOS()
	default:
		return ""
	}
}

/* -------------------- Unexported Functions -------------------- */

func wifiInfo() string {
	cmd := exec.Command(osxWifiCmd, osxWifiArg)
	return wtf.ExecuteCommand(cmd)
}

func matchStr(data [][]string) string {
	if len(data) <= 1 {
		return ""
	} else {
		return data[1][1]
	}
}
