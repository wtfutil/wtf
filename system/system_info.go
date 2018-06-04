package system

import (
	"github.com/matishsiao/goInfo"
)

//systemInfo will hold system information using goInfo pkg
type SystemInfo struct {
	GoOs     string
	Kernel   string
	Version  string
	Platform string
	OS       string
	Hostname string
	CPUs     int
}

//NewSystemInfo will get  current system info, should work on linux/mac/win but I could only test it on linux
func NewSystemInfo() *SystemInfo {
	gi := goInfo.GetInfo()
	var sysInfo *SystemInfo
	sysInfo = &SystemInfo{
		GoOs:     gi.GoOS,
		Kernel:   gi.Kernel,
		Version:  gi.Core,
		Platform: gi.Platform,
		OS:       gi.OS,
		Hostname: gi.Hostname,
		CPUs:     gi.CPUs,
	}
	return sysInfo
}
