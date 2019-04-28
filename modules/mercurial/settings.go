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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		commitCount:  ymlConfig.UInt("commitCount", 10),
		commitFormat: ymlConfig.UString("commitFormat", "[forestgreen]{rev}:{phase} [white]{desc|firstline|strip} [grey]{author|person} {date|age}[white]"),
		repositories: ymlConfig.UList("repositories"),
	}

	return &settings
}
