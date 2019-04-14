package git

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	commitCount  int
	commitFormat string
	dateFormat   string
	repositories []interface{}
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.git")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		commitCount:  localConfig.UInt("commitCount", 10),
		commitFormat: localConfig.UString("commitFormat", "[forestgreen]%h [white]%s [grey]%an on %cd[white]"),
		dateFormat:   localConfig.UString("dateFormat", "%b %d, %Y"),
		repositories: localConfig.UList("repositories"),
	}

	return &settings
}
