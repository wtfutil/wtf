package security

import (
	"io/ioutil"
	"os/exec"
	"regexp"
)

// https://github.com/yelinaung/wifi-name/blob/master/wifi-name.go
const osxWifiCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"

/* -------------------- Exported Functions -------------------- */

func WifiEncryption() string {
	cmd := exec.Command(osxWifiCmd, "-I")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return ""
	}

	if err := cmd.Start(); err != nil {
		return ""
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	name := findMatch(`s*auth: (.+)s*`, str)
	return matchStr(name)
}

func WifiName() string {
	cmd := exec.Command(osxWifiCmd, "-I")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return ""
	}

	if err := cmd.Start(); err != nil {
		return ""
	}

	var str string
	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	name := findMatch(`s*SSID: (.+)s*`, str)
	return matchStr(name)
}

/* -------------------- Unexported Functions -------------------- */

func findMatch(pattern string, data string) [][]string {
	r := regexp.MustCompile(pattern)

	name := r.FindAllStringSubmatch(data, -1)
	return name
}

func matchStr(data [][]string) string {
	if len(data) <= 1 {
		return ""
	} else {
		return data[1][1]
	}
}
