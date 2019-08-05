package utils

// OpenFileUtil defines the system utility to use to open files
var OpenFileUtil = "open"

// Init initializes global settings in the wtf package
func Init(openFileUtil string) {
	OpenFileUtil = openFileUtil
}
