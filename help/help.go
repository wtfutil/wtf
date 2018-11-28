package help

import (
	"fmt"

	"github.com/wtfutil/wtf/git"
	"github.com/wtfutil/wtf/github"
	"github.com/wtfutil/wtf/textfile"
	"github.com/wtfutil/wtf/todo"
	"github.com/wtfutil/wtf/todoist"
	"github.com/wtfutil/wtf/weatherservices/weather"
)

func Display(moduleName string) {
	if moduleName == "" {
		fmt.Println("\n  --module takes a module name as an argument, i.e: '--module=github'")
	} else {
		fmt.Printf("%s\n", helpFor(moduleName))
	}
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
	case "todoist":
		return todoist.HelpText
	case "weather":
		return weather.HelpText
	default:
		return fmt.Sprintf("\n  There is no help available for '%s'", moduleName)
	}
}
