package gspreadsheets

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "Google Spreadsheets"
)

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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		cellNames:  ymlConfig.UList("cells.names"),
		secretFile: ymlConfig.UString("secretFile"),
		sheetID:    ymlConfig.UString("sheetId"),
	}

	settings.colors.values = ymlConfig.UString("colors.values", "green")

	return &settings
}
