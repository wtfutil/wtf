package cmdrunner

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const configKey = "cmdrunner"

type Settings struct {
	common *cfg.Common

	args []string
	cmd  string
}

func NewSettingsFromYAML(name string, moduleConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, moduleConfig, globalConfig),

		args: wtf.ToStrs(moduleConfig.UList("args")),
		cmd:  moduleConfig.UString("cmd"),
	}

	return &settings
}
