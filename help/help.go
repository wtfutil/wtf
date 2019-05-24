package help

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/maker"
	"github.com/wtfutil/wtf/utils"
)

func Display(moduleName string, config *config.Config) {
	if moduleName == "" {
		fmt.Println("\n  --module takes a module name as an argument, i.e: '--module=github'")
	} else {
		fmt.Printf("%s\n", helpFor(moduleName, config))
	}
}

func helpFor(moduleName string, config *config.Config) string {
	widget := maker.MakeWidget(nil, nil, moduleName, moduleName, config, config)
	result := ""
	result += utils.StripColorTags(widget.HelpText())
	result += widget.ConfigText()
	return result
}
