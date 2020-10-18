package finnhub

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
	"os"
)

const (
	defaultFocusable = true
	defaultTitle     = "Stocks"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey  string   `help:"Your finnhub API token."`
	symbols []string `help:"An array of stocks symbols (i.e. AAPL, MSFT)"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:  ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_FINNHUB_API_KEY"))),
		symbols: utils.ToStrs(ymlConfig.UList("symbols")),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
