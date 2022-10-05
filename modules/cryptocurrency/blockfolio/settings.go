package blockfolio

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Blockfolio"
)

type colors struct {
	name  string
	grows string
	drop  string
}

type Settings struct {
	*cfg.Common

	colors

	deviceToken     string
	displayHoldings bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		deviceToken:     ymlConfig.UString("device_token"),
		displayHoldings: ymlConfig.UBool("displayHoldings", true),
	}

	settings.SetDocumentationPath("cryptocurrencies/blockfolio")

	return &settings
}
