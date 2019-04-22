package prettyweather

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "prettyweather"

type Settings struct {
	common *cfg.Common

	city     string
	unit     string
	view     string
	language string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		city:     localConfig.UString("city", "Barcelona"),
		language: localConfig.UString("language", "en"),
		unit:     localConfig.UString("unit", "m"),
		view:     localConfig.UString("view", "0"),
	}

	return &settings
}
