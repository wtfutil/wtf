package prettyweather

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	city     string
	unit     string
	view     string
	language string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.prettyweather")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		city:     localConfig.UString("city", "Barcelona"),
		language: localConfig.UString("language", "en"),
		unit:     localConfig.UString("unit", "m"),
		view:     localConfig.UString("view", "0"),
	}

	return &settings
}
