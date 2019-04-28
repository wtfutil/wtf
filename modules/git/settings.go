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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		commitCount:  ymlConfig.UInt("commitCount", 10),
		commitFormat: ymlConfig.UString("commitFormat", "[forestgreen]%h [white]%s [grey]%an on %cd[white]"),
		dateFormat:   ymlConfig.UString("dateFormat", "%b %d, %Y"),
		repositories: ymlConfig.UList("repositories"),
	}

	return &settings
}
