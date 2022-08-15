package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

// https://github.com/yelinaung/wifi-name/blob/master/wifi-name.go
const osxWifiCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxWifiArg = "-I"

/* -------------------- Exported Functions -------------------- */

func WifiEncryption() string {
	switch runtime.GOOS {
	case "linux":
		return wifiEncryptionLinux()
	case "darwin":
		return wifiEncryptionMacOS()
	case "windows":
		return wifiEncryptionWindows()
	default:
		return ""
	}
}

func WifiName() string {
	switch runtime.GOOS {
	case "linux":
		return wifiNameLinux()
	case "darwin":
		return wifiNameMacOS()
	case "windows":
		return wifiNameWindows()
	default:
		return ""
	}
}

/* -------------------- Unexported Functions -------------------- */

func wifiEncryptionLinux() string {
	cmd := exec.Command("nmcli", "-t", "-f", "in-use,security", "dev", "wifi")
	out := utils.ExecuteCommand(cmd)

	name := utils.FindMatch(`\*:(.+)`, out)

	if len(name) > 0 {
		return name[0][1]
	}

	return ""
}

func wifiEncryptionMacOS() string {
	name := utils.FindMatch(`s*auth: (.+)s*`, wifiInfo())
	return matchStr(name)
}

func wifiInfo() string {
	cmd := exec.Command(osxWifiCmd, osxWifiArg)
	return utils.ExecuteCommand(cmd)
}

func wifiNameLinux() string {
	cmd, _ := exec.Command("iwgetid", "-r").Output()
	return string(cmd)
}

func wifiNameMacOS() string {
	name := utils.FindMatch(`s*SSID: (.+)s*`, wifiInfo())
	return matchStr(name)
}

func matchStr(data [][]string) string {
	if len(data) <= 1 {
		return ""
	}

	return data[1][1]
}

// Windows
func wifiEncryptionWindows() string {
	return parseWlanNetsh("Authentication")
}

func wifiNameWindows() string {
	return parseWlanNetsh("SSID")
}

func parseWlanNetsh(target string) string {
	cmd := exec.Command("netsh.exe", "wlan", "show", "interfaces")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	splits := strings.Split(string(out), "\n")
	var words []string
	for _, line := range splits {
		token := strings.Split(line, ":")
		for _, word := range token {
			words = append(words, strings.TrimSpace(word))
		}
	}
	for i, token := range words {
		if token == target {
			return words[i+1]
		}
	}
	return "N/A"
}
