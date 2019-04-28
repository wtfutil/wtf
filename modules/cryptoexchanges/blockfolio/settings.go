package blockfolio

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "blockfolio"

type colors struct {
	name  string
	grows string
	drop  string
}

type Settings struct {
	colors
	common *cfg.Common

	deviceToken     string
	displayHoldings bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		deviceToken:     ymlConfig.UString("device_token"),
		displayHoldings: ymlConfig.UBool("displayHoldings", true),
	}

	return &settings
}
