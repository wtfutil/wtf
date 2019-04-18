package ipinfo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "ipinfo"

type colors struct {
	name  string
	value string
}

type Settings struct {
	colors
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),
	}

	settings.colors.name = localConfig.UString("colors.name", "red")
	settings.colors.value = localConfig.UString("colors.value", "white")

	return &settings
}
