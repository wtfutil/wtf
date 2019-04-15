package ipapi

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type colors struct {
	name  string
	value string
}

type Settings struct {
	colors
	common *cfg.Common
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.ipapi")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),
	}

	settings.colors.name = localConfig.UString("colors.name", "red")
	settings.colors.value = localConfig.UString("colors.value", "white")

	return &settings
}
