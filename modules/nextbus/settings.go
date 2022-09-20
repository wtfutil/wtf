package nextbus

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "nextbus"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	route  string `help:"Route Number of your bus"`
	agency string `help:"Transit agency of your bus"`
	stopID string `help:"Your bus stop number"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		route:  ymlConfig.UString("route"),
		agency: ymlConfig.UString("agency"),
		stopID: ymlConfig.UString("stopID"),
	}

	return &settings
}
