package bittrex

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Bittrex"
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

type currency struct {
	displayName string
	market      []interface{}
}

type summary struct {
	currencies map[string]*currency
}

type Settings struct {
	*cfg.Common

	colors
	summary
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.base.name = ymlConfig.UString("colors.base.name")
	settings.base.displayName = ymlConfig.UString("colors.base.displayName")

	settings.market.name = ymlConfig.UString("colors.market.name")
	settings.market.field = ymlConfig.UString("colors.market.field")
	settings.market.value = ymlConfig.UString("colors.market.value")

	settings.currencies = make(map[string]*currency)
	for key, val := range ymlConfig.UMap("summary") {
		coercedVal := val.(map[string]interface{})

		currency := &currency{
			displayName: coercedVal["displayName"].(string),
			market:      coercedVal["market"].([]interface{}),
		}

		settings.currencies[key] = currency
	}

	settings.SetDocumentationPath("cryptocurrencies/bittrex")

	return &settings
}
