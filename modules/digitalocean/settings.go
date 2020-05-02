package digitalocean

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const (
	defaultFocusable = true
	defaultTitle     = "DigitalOcean"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey     string `help:"Your DigitalOcean API key."`
	dateFormat string `help:"The format to display dates and times in."`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:     ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_DIGITALOCEAN_API_KEY"))),
		dateFormat: ymlConfig.UString("dateFormat", wtf.DateFormat),
	}

	cfg.ConfigureSecret(
		globalConfig,
		"",
		name,
		nil,
		&settings.apiKey,
	)

	return &settings
}
