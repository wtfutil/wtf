package gspreadsheets

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "gspreadsheets"

type colors struct {
	values string
}

type Settings struct {
	colors
	common *cfg.Common

	cellAddresses []interface{}
	cellNames     []interface{}
	secretFile    string
	sheetID       string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		cellNames:  localConfig.UList("cells.names"),
		secretFile: localConfig.UString("secretFile"),
		sheetID:    localConfig.UString("sheetId"),
	}

	settings.colors.values = localConfig.UString("colors.values", "green")

	return &settings
}
