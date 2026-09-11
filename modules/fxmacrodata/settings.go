package fxmacrodata

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "Macro Releases"
	defaultBaseURL   = "https://api.fxmacrodata.com"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey     string
	baseURL    string
	currencies []string `help:"An array of currency codes to show releases for (for example: USD, EUR)."`
	count      int      `help:"The number of upcoming releases to display."`
	topTier    bool     `help:"When true, show only the releases marked top tier for their currency."`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.common.RefreshInterval = cfg.ParseTimeString(ymlConfig, "refreshInterval", "15m")

	// USD data is public, so an absent key is an ordinary configuration rather
	// than an error. A key widens the module to the other currencies.
	settings.apiKey = ymlConfig.UString("apiKey", os.Getenv("FXMACRODATA_API_KEY"))
	settings.baseURL = ymlConfig.UString("baseURL", defaultBaseURL)
	settings.count = ymlConfig.UInt("count", 10)
	settings.topTier = ymlConfig.UBool("topTier", false)

	settings.currencies = utils.ToStrs(ymlConfig.UList("currencies"))
	if len(settings.currencies) == 0 {
		settings.currencies = []string{"USD"}
	}

	return &settings
}
