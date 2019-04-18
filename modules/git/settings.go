package git

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "git"

type Settings struct {
	common *cfg.Common

	commitCount  int
	commitFormat string
	dateFormat   string
	repositories []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		commitCount:  localConfig.UInt("commitCount", 10),
		commitFormat: localConfig.UString("commitFormat", "[forestgreen]%h [white]%s [grey]%an on %cd[white]"),
		dateFormat:   localConfig.UString("dateFormat", "%b %d, %Y"),
		repositories: localConfig.UList("repositories"),
	}

	return &settings
}
