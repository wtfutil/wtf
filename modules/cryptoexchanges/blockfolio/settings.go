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

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		deviceToken:     localConfig.UString("device_token"),
		displayHoldings: localConfig.UBool("displayHoldings", true),
	}

	return &settings
}
