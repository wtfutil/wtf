package spotify

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Spotify"
)

type colors struct {
	label string
	text  string
}

type Settings struct {
	colors
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.colors.label = ymlConfig.UString("colors.label", "green")
	settings.colors.text = ymlConfig.UString("colors.text", "white")

	return &settings
}
