// Package exchangerates
package exchangerates

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Exchange rates"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	rates map[string][]string `help:"Defines what currency rates we want to know about`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		rates: map[string][]string{},
	}

	raw := ymlConfig.UMap("rates", map[string]interface{}{})
	for key, value := range raw {
		settings.rates[key] = []string{}
		switch value.(type) {
		case string:
			settings.rates[key] = []string{value.(string)}
		case []interface{}:
			for _, currency := range value.([]interface{}) {
				str, ok := currency.(string)
				if ok {
					settings.rates[key] = append(settings.rates[key], str)
				}
			}
		}
	}

	return &settings
}
