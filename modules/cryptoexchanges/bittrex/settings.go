package bittrex

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type colors struct {
	base struct {
		name        string
		displayName string
	}
	market struct {
		name  string
		field string
		value string
	}
}

type Settings struct {
	colors
	common *cfg.Common

	summary map[string]interface{}
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.bittrex")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),
	}

	settings.colors.base.name = localConfig.UString("colors.base.name")
	settings.colors.base.displayName = localConfig.UString("colors.base.displayName")

	settings.colors.market.name = localConfig.UString("colors.market.name")
	settings.colors.market.field = localConfig.UString("colors.market.field")
	settings.colors.market.value = localConfig.UString("colors.market.value")

	summaryMap, _ := ymlConfig.Map("summary")
	settings.summary = summaryMap

	return &settings
}
