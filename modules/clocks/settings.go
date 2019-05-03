package clocks

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

const defaultTitle = "Clocks"

type colors struct {
	rows struct {
		even string
		odd  string
	}
}

type Settings struct {
	colors
	common *cfg.Common

	dateFormat string
	timeFormat string
	locations  map[string]interface{}
	sort       string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		dateFormat: ymlConfig.UString("dateFormat", wtf.SimpleDateFormat),
		timeFormat: ymlConfig.UString("timeFormat", wtf.SimpleTimeFormat),
		locations:  ymlConfig.UMap("locations"),
		sort:       ymlConfig.UString("sort"),
	}

	settings.colors.rows.even = ymlConfig.UString("colors.rows.even", "white")
	settings.colors.rows.odd = ymlConfig.UString("colors.rows.odd", "blue")

	return &settings
}
