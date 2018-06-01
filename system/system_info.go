package system

import (
	"os/exec"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type SystemInfo struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

func NewSystemInfo() *SystemInfo {
	m := make(map[string]string)

	arg := []string{}

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("lsb_release -a", arg...)
	case "darwin":
		cmd := exec.Command("sw_vers", arg...)
	default:
		return ""
	}
	
	raw := wtf.ExecuteCommand(cmd)

	for _, row := range strings.Split(raw, "\n") {
		parts := strings.Split(row, ":")

		if len(parts) < 2 {
			continue
		}

		m[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	sysInfo := SystemInfo{
		ProductName:    m["ProductName"],
		ProductVersion: m["ProductVersion"],
		BuildVersion:   m["BuildVersion"],
	}

	return &sysInfo
}
