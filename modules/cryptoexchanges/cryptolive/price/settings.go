package price

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "cryptolive"

type colors struct {
	from struct {
		name        string
		displayName string
	}
	to struct {
		name  string
		price string
	}
	top struct {
		from struct {
			name        string
			displayName string
		}
		to struct {
			name  string
			field string
			value string
		}
	}
}

type currency struct {
	displayName string
	to          []interface{}
}

type Settings struct {
	colors
	common     *cfg.Common
	currencies map[string]*currency
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),
	}

	settings.colors.from.name = localConfig.UString("colors.from.name")
	settings.colors.from.displayName = localConfig.UString("colors.from.displayName")

	settings.colors.to.name = localConfig.UString("colors.to.name")
	settings.colors.to.price = localConfig.UString("colors.to.price")

	settings.colors.top.from.name = localConfig.UString("colors.top.from.name")
	settings.colors.top.from.displayName = localConfig.UString("colors.top.from.displayName")

	settings.colors.top.to.name = localConfig.UString("colors.top.to.name")
	settings.colors.top.to.field = localConfig.UString("colors.top.to.field")
	settings.colors.top.to.value = localConfig.UString("colors.top.to.value")

	settings.currencies = make(map[string]*currency)

	for key, val := range localConfig.UMap("currencies") {
		coercedVal := val.(map[string]interface{})

		currency := &currency{
			displayName: coercedVal["displayName"].(string),
			to:          coercedVal["to"].([]interface{}),
		}

		settings.currencies[key] = currency
	}

	return &settings
}
