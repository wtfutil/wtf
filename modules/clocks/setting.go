package clocks

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

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

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.clocks")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		dateFormat: localConfig.UString("dateFormat", wtf.SimpleDateFormat),
		timeFormat: localConfig.UString("timeFormat", wtf.SimpleTimeFormat),
		locations:  localConfig.UMap("locations"),
		sort:       localConfig.UString("sort"),
	}

	settings.colors.rows.even = localConfig.UString("colors.rows.even", "white")
	settings.colors.rows.odd = localConfig.UString("colors.rows.odd", "blue")

	return &settings
}
