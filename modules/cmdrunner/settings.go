package cmdrunner

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

type Settings struct {
	common *cfg.Common

	args []string
	cmd  string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.cmdrunner")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		args: wtf.ToStrs(localConfig.UList("args")),
		cmd:  localConfig.UString("cmd"),
	}

	return &settings
}
