package digitalclock

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Clocks"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	hourFormat string `help:"The format of the clock." values:"12 or 24"`
	color      string `help:"The color of the clock."`
	font       string `help:"The font of the clock." values:"bigfont or digitalfont"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		color:      ymlConfig.UString("color"),
		font:       ymlConfig.UString("font"),
		hourFormat: ymlConfig.UString("hourFormat", "24"),
	}

	return &settings
}
