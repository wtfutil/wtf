package uptimekuma

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Uptime Kuma"
)

type Settings struct {
	common *cfg.Common

	url string `help:"Status page URL; e.g. https://uptimekuma.example.com/status/overview"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		// Configure your settings attributes here. See http://github.com/olebedev/config for type details
		url: ymlConfig.UString("url"),
	}

	return &settings
}
