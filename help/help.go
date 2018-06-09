package help

import (
	"fmt"
	"os"

	"github.com/andrewzolotukhin/wtf/git"
	"github.com/andrewzolotukhin/wtf/github"
	"github.com/andrewzolotukhin/wtf/textfile"
	"github.com/andrewzolotukhin/wtf/todo"
	"github.com/andrewzolotukhin/wtf/weather"
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
