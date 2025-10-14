package azurelogs

import (
	"github.com/olebedev/config"

	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Azure Logs"
)

// Settings defines the configuration for the Azure Logs widget
type Settings struct {
	*cfg.Common

	// Queryfile is the path to the YAML file containing the Azure query configuration
	Queryfile string `help:"Path to YAML file containing Azure Log Analytics query configuration"`
}

// NewSettingsFromYAML creates a new Settings instance from YAML configuration
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		Queryfile: ymlConfig.UString("queryFile", ""),
	}

	return &settings
}
