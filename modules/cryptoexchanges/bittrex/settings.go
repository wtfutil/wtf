package bittrex

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "bittrex"

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

type currency struct {
	displayName string
	market      []interface{}
}

type summary struct {
	currencies map[string]*currency
}

type Settings struct {
	colors
	common *cfg.Common
	summary
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),
	}

	settings.colors.base.name = localConfig.UString("colors.base.name")
	settings.colors.base.displayName = localConfig.UString("colors.base.displayName")

	settings.colors.market.name = localConfig.UString("colors.market.name")
	settings.colors.market.field = localConfig.UString("colors.market.field")
	settings.colors.market.value = localConfig.UString("colors.market.value")

	settings.summary.currencies = make(map[string]*currency)
	for key, val := range localConfig.UMap("summary") {
		coercedVal := val.(map[string]interface{})

		currency := &currency{
			displayName: coercedVal["displayName"].(string),
			market:      coercedVal["market"].([]interface{}),
		}

		settings.currencies[key] = currency
	}

	return &settings
}
