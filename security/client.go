package security

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	//"runtime"
	//"strings"
)

const fwGlobalState = "/usr/libexec/ApplicationFirewall/socketfilterfw --getglobalstate"
const fwStealthMode = "/usr/libexec/ApplicationFirewall/socketfilterfw --getstealthmode"

// https://github.com/yelinaung/wifi-name/blob/master/wifi-name.go
const osxCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxArgs = "-I"

func Fetch() map[string]string {
	data := make(map[string]string)

	data["Wifi"] = WifiName()

	return data
}

func WifiName() string {
	cmd := exec.Command(osxCmd, osxArgs)

	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	r := regexp.MustCompile(`s*SSID: (.+)s*`)

	name := r.FindAllStringSubmatch(str, -1)

	if len(name) <= 1 {
		return ""
	} else {
		return name[1][1]
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
