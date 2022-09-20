package toplist

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
	limit       int
	to          []interface{}
}

type Settings struct {
	*cfg.Common

	colors
	top map[string]*currency
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}

	settings.colors.from.name = ymlConfig.UString("colors.from.name")
	settings.colors.from.displayName = ymlConfig.UString("colors.from.displayName")

	settings.colors.to.name = ymlConfig.UString("colors.to.name")
	settings.colors.to.price = ymlConfig.UString("colors.to.price")

	settings.colors.top.from.name = ymlConfig.UString("colors.top.from.name")
	settings.colors.top.from.displayName = ymlConfig.UString("colors.top.from.displayName")

	settings.colors.top.to.name = ymlConfig.UString("colors.top.to.name")
	settings.colors.top.to.field = ymlConfig.UString("colors.top.to.field")
	settings.colors.top.to.value = ymlConfig.UString("colors.top.to.value")

	settings.top = make(map[string]*currency)

	for key, val := range ymlConfig.UMap("top") {
		coercedVal := val.(map[string]interface{})

		limit, _ := coercedVal["limit"].(int)

		currency := &currency{
			displayName: coercedVal["displayName"].(string),
			limit:       limit,
			to:          coercedVal["to"].([]interface{}),
		}

		settings.top[key] = currency
	}

	return &settings
}
