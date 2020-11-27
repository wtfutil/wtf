package prettyweather

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Pretty Weather"
)

type Settings struct {
	*cfg.Common

	city     string
	unit     string
	view     string
	language string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		city:     ymlConfig.UString("city", "Barcelona"),
		language: ymlConfig.UString("language", "en"),
		unit:     ymlConfig.UString("unit", "m"),
		view:     ymlConfig.UString("view", "0"),
	}

	settings.SetDocumentationPath("weather_services/prettyweather")

	return &settings
}
