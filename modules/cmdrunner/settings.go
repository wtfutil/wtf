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

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		args: wtf.ToStrs(localConfig.UList("args")),
		cmd:  localConfig.UString("cmd"),
	}

	return &settings
}
