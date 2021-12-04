//go:build !windows

package system

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

type SystemInfo struct {
	ProductName    string
	ProductVersion string
	BuildVersion   string
}

func NewSystemInfo() *SystemInfo {
	m := make(map[string]string)

	arg := []string{}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		arg = append(arg, "-a")
		cmd = exec.Command("lsb_release", arg...)
	case "darwin":
		cmd = exec.Command("sw_vers", arg...)
	default:
		cmd = exec.Command("sw_vers", arg...)
	}

	raw := utils.ExecuteCommand(cmd)

	for _, row := range strings.Split(raw, "\n") {
		parts := strings.Split(row, ":")
		if len(parts) < 2 {
			continue
		}

		m[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	var sysInfo *SystemInfo
	switch runtime.GOOS {
	case "linux":
		sysInfo = &SystemInfo{
			ProductName:    m["Distributor ID"],
			ProductVersion: m["Description"],
			BuildVersion:   m["Release"],
		}
	default:
		sysInfo = &SystemInfo{
			ProductName:    m["ProductName"],
			ProductVersion: m["ProductVersion"],
			BuildVersion:   m["BuildVersion"],
		}

	}
	return sysInfo
}
