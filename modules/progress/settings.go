package progress

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Progress"
)

type colors struct {
	gradientA string `help:"Start color for linear gradient." values:"any X11 or hex color" optional:"true" default:"#56ab2f"`
	gradientB string `help:"End color for linear gradient." values:"any X11 or hex color" optional:"true" default:"#a8e063"`
	solid     string `help:"Use a solid color instead of linear color gradient ." values:"any X11 or hex color" optional:"true"`
}

// Settings defines the configuration properties for this module
type Settings struct {
	colors
	common *cfg.Common

	showPercentage string `help:"Where to display the percentage" values:"left, right, above, below, titleLeft, titleRight, or none" optional:"true" default:"right"`
	padding        int    `help:"Amount of spaces to add as left/right padding." values:"A positive integer, 0..n" optional:"true" default:"1"`

	minimum float64 `help:"Minimum progress value." values:"A positive decimal value, 0.0..n.n" optional:"true" default:"0"`
	maximum float64 `help:"Maximum progress value." values:"A positive decimal value, 0.0..n.n" optional:"true" default:"0"`
	current float64 `help:"Current progress value. If maximum value is 0, current value is assumed to be a percentage between 0-100." values:"A positive decimal value, 0.0..n.n" optional:"true" default:"0"`

	minimumCmd string `help:"Execute shell command to determine minimum progress value. Return value must be numeric." values:"Any shell command" optional:"true"`
	maximumCmd string `help:"Execute shell command to determine maximum progress value. Return value must be numeric." values:"Any shell command" optional:"true"`
	currentCmd string `help:"Execute shell command to determine current progress value. Return value must be numeric." values:"Any shell command" optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		showPercentage: ymlConfig.UString("showPercentage", "right"),
		padding:        ymlConfig.UInt("padding", 1),
		minimum:        ymlConfig.UFloat64("minimum", 0),
		maximum:        ymlConfig.UFloat64("maximum", 0),
		current:        ymlConfig.UFloat64("current", 0),
		minimumCmd:     ymlConfig.UString("minimumCmd", ""),
		maximumCmd:     ymlConfig.UString("maximumCmd", ""),
		currentCmd:     ymlConfig.UString("currentCmd", ""),
	}

	settings.colors.gradientA = ymlConfig.UString("colors.gradientA", "#56ab2f")
	settings.colors.gradientB = ymlConfig.UString("colors.gradientB", "#a8e063")
	settings.colors.solid = ymlConfig.UString("colors.solid", "")

	return &settings
}
