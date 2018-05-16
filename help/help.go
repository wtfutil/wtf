package help

import (
	"fmt"
	"os"
)

func DisplayCommandInfo(args []string, version string) {
	if len(args) == 0 {
		return
	}

	cmd := args[0]

	switch cmd {
	case "help", "--help":
		DisplayHelpInfo()
	case "version", "--version":
		DisplayVersionInfo(version)
	}

}

func DisplayHelpInfo() {
	fmt.Println("Help is coming...")
	os.Exit(0)
}

func DisplayVersionInfo(version string) {
	fmt.Printf("Version: %s\n", version)
	os.Exit(0)
}
