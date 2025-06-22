package price

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "CryptoLive"
)

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
	*cfg.Common

	colors
	currencies map[string]*currency
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.from.name = ymlConfig.UString("colors.from.name")
	settings.from.displayName = ymlConfig.UString("colors.from.displayName")

	settings.to.name = ymlConfig.UString("colors.to.name")
	settings.to.price = ymlConfig.UString("colors.to.price")

	settings.top.from.name = ymlConfig.UString("colors.top.from.name")
	settings.top.from.displayName = ymlConfig.UString("colors.top.from.displayName")

	settings.top.to.name = ymlConfig.UString("colors.top.to.name")
	settings.top.to.field = ymlConfig.UString("colors.top.to.field")
	settings.top.to.value = ymlConfig.UString("colors.top.to.value")

	settings.currencies = make(map[string]*currency)

	for key, val := range ymlConfig.UMap("currencies") {
		coercedVal := val.(map[string]interface{})

		currency := &currency{
			displayName: coercedVal["displayName"].(string),
			to:          coercedVal["to"].([]interface{}),
		}

		settings.currencies[key] = currency
	}

	return &settings
}
