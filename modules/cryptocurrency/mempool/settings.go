package mempool

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "mempool"
	defaultAPIURL    = "https://mempool.space/api/v1/fees/recommended"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiURL string `help:"Fee recommendation endpoint URL." optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiURL: ymlConfig.UString("apiURL", defaultAPIURL),
	}

	return &settings
}

func (widget *Widget) ConfigText() string {
	return utils.HelpFromInterface(Settings{})
}
