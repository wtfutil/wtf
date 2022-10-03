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
	*cfg.Common

	color          string `help:"The color of the clock."`
	font           string `help:"The font of the clock." values:"bigfont or digitalfont"`
	hourFormat     string `help:"The format of the clock." values:"12 or 24"`
	dateFormat     string `help:"The format of the date."`
	withDate       bool   `help:"Whether or not to display date information"`
	withDatePrefix bool   `help:"Whether or not to display Date: prefix"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		color:          ymlConfig.UString("color"),
		font:           ymlConfig.UString("font"),
		hourFormat:     ymlConfig.UString("hourFormat", "24"),
		dateFormat:     ymlConfig.UString("dateFormat", "Monday January 02 2006"),
		withDate:       ymlConfig.UBool("withDate", true),
		withDatePrefix: ymlConfig.UBool("withDatePrefix", true),
	}

	return &settings
}
