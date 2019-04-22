package mercurial

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "mercurial"

type Settings struct {
	common *cfg.Common

	commitCount  int
	commitFormat string
	repositories []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		commitCount:  localConfig.UInt("commitCount", 10),
		commitFormat: localConfig.UString("commitFormat", "[forestgreen]{rev}:{phase} [white]{desc|firstline|strip} [grey]{author|person} {date|age}[white]"),
		repositories: localConfig.UList("repositories"),
	}

	return &settings
}
