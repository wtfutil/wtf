package security

import (
	"os/exec"

	"github.com/senorprogrammer/wtf/wtf"
)

// https://github.com/yelinaung/wifi-name/blob/master/wifi-name.go
const osxWifiCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxWifiArg = "-I"

/* -------------------- Exported Functions -------------------- */

func WifiEncryption() string {
	name := wtf.FindMatch(`s*auth: (.+)s*`, wifiInfo())
	return matchStr(name)
}

func WifiName() string {
	name := wtf.FindMatch(`s*SSID: (.+)s*`, wifiInfo())
	return matchStr(name)
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
