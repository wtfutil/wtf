package yfinance

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "Yahoo Finance"
)

type colors struct {
	bigup   string
	up      string
	drop    string
	bigdrop string
}

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	colors  colors
	sort    bool
	symbols []string `help:"An array of Yahoo Finance symbols (for example: DOCN, GME, GC=F)"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		// RefreshInterval: ,
	}

	settings.common.RefreshInterval = cfg.ParseTimeString(ymlConfig, "refreshInterval", "60s")
	settings.colors.bigup = ymlConfig.UString("colors.bigup", "greenyellow")
	settings.colors.up = ymlConfig.UString("colors.up", "green")
	settings.colors.drop = ymlConfig.UString("colors.drop", "firebrick")
	settings.colors.bigdrop = ymlConfig.UString("colors.bigdrop", "red")
	settings.sort = ymlConfig.UBool("sort", false)
	settings.symbols = utils.ToStrs(ymlConfig.UList("symbols"))
	return &settings
}
