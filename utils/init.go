package utils

// OpenFileUtil defines the system utility to use to open files
var OpenFileUtil = "open"
var OpenUrlUtil = []string{}

// Init initializes global settings in the wtf package
func Init(openFileUtil string, openUrlUtil []string) {
	OpenFileUtil = openFileUtil
	OpenUrlUtil = openUrlUtil
}
