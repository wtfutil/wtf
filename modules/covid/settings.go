package covid

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Covid tracker"
)

// Settings is the struct for this module's settings
type Settings struct {
	*cfg.Common

	country string `help:"Country (code) from which to retrieve stats."`
}

// NewSettingsFromYAML returns the settings from the config yaml file
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		country: ymlConfig.UString("country"),
	}

	return &settings
}
