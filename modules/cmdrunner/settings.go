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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		args: wtf.ToStrs(ymlConfig.UList("args")),
		cmd:  ymlConfig.UString("cmd"),
	}

	return &settings
}
