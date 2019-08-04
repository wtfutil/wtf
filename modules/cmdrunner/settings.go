package cmdrunner

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const defaultTitle = "CmdRunner"

type Settings struct {
	common *cfg.Common

	args []string `help:"The arguments to the command, with each item as an element in an array. Example: for curl -I cisco.com, the arguments array would be ['-I', 'cisco.com']."`
	cmd  string   `help:"The terminal command to be run, withouth the arguments. Ie: ping, whoami, curl."`
}

func NewSettingsFromYAML(name string, moduleConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, moduleConfig, globalConfig),

		args: wtf.ToStrs(moduleConfig.UList("args")),
		cmd:  moduleConfig.UString("cmd"),
	}

	return &settings
}
