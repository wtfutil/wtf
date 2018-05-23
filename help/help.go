package help

import (
	"fmt"
	"os"

	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/textfile"
	"github.com/senorprogrammer/wtf/todo"
	"github.com/senorprogrammer/wtf/weather"
)

func DisplayCommandInfo(args []string, version string) {
	if len(args) == 0 {
		return
	}

	cmd := args[0]

	switch cmd {
	case "help", "--help":
		DisplayHelpInfo(args)
	case "version", "--version":
		DisplayVersionInfo(version)
	}

}

func DisplayHelpInfo(args []string) {
	if len(args) >= 1 {
		fmt.Printf("%s\n", helpFor(args[0]))
	} else {
		fmt.Println("\n  --help takes a module name as an argument, i.e: '--help github'")
	}

	os.Exit(0)
}

func DisplayVersionInfo(version string) {
	fmt.Printf("Version: %s\n", version)
	os.Exit(0)
}

func helpFor(moduleName string) string {
	switch moduleName {
	case "git":
		return git.HelpText
	case "github":
		return github.HelpText
	case "textfile":
		return textfile.HelpText
	case "todo":
		return todo.HelpText
	case "weather":
		return weather.HelpText
	default:
		return fmt.Sprintf("\n  There is no help available for '%s'", moduleName)
	}
}
