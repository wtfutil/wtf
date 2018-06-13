package help

import (
	"fmt"
	"os"

	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/textfile"
	"github.com/senorprogrammer/wtf/todo"
	"github.com/senorprogrammer/wtf/weatherservices/weather"
)

func DisplayModuleInfo(moduleName string) {
	if moduleName != "" {
		fmt.Printf("%s\n", helpFor(moduleName))
	} else {
		fmt.Println("\n  --module takes a module name as an argument, i.e: '--module=github'")
	}

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
