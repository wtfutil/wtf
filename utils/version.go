package utils

import "runtime/debug"

// VersionTime returns version and release time of executable.
func VersionTime() (string, string) {
	info, ok := debug.ReadBuildInfo()
	var settings []debug.BuildSetting
	if !ok {
		settings = []debug.BuildSetting{}
	} else {
		settings = info.Settings
	}

	return extractVersionTime(settings)
}

func extractVersionTime(settings []debug.BuildSetting) (string, string) {
	version := "dev"
	time := "now"
	for _, setting := range settings {
		if setting.Key == "vcs.revision" {
			version = setting.Value
		} else if setting.Key == "vcs.time" {
			time = setting.Value
		}
	}
	return version, time
}
