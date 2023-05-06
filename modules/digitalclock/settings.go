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
	dateTitle      bool   `help:"Whether or not to display date as widget title"`
	withDate       bool   `help:"Whether or not to display date information"`
	withUTC        bool   `help:"Whether or not to display UTC information"`
	withEpoch      bool   `help:"Whether or not to display Epoch information"`
	withDatePrefix bool   `help:"Whether or not to display Date: prefix"`
	centerAlign    bool   `help:"Whether or not to use center align in widget"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		color:          ymlConfig.UString("color"),
		font:           ymlConfig.UString("font"),
		hourFormat:     ymlConfig.UString("hourFormat", "24"),
		dateFormat:     ymlConfig.UString("dateFormat", "Monday January 02 2006"),
		dateTitle:      ymlConfig.UBool("dateTitle", false),
		withDate:       ymlConfig.UBool("withDate", true),
		withUTC:        ymlConfig.UBool("withUTC", true),
		withEpoch:      ymlConfig.UBool("withEpoch", true),
		withDatePrefix: ymlConfig.UBool("withDatePrefix", true),
		centerAlign:    ymlConfig.UBool("centerAlign", false),
	}

	return &settings
}
