package todo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Todo"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	filePath   string
	checked    string
	unchecked  string
	newPos     string
	checkedPos string
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	common := cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig)

	settings := Settings{
		Common: common,

		filePath:  ymlConfig.UString("filename"),
		checked:   ymlConfig.UString("checkedIcon", common.Checkbox.Checked),
		unchecked: ymlConfig.UString("uncheckedIcon", common.Checkbox.Unchecked),
		newPos: ymlConfig.UString("newPos", "first"),
		checkedPos: ymlConfig.UString("checkedPos", "last"),
	}

	switch settings.newPos {
	case "first","last":
	default:
		settings.newPos = "last"
	}
	switch settings.checkedPos {
	case "first", "last", "none":
	default:
		settings.checkedPos = "last"
	}

	return &settings
}
