package textfile

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Textfile"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	filePaths   []interface{}
	format      bool
	formatStyle string
	wrapText    bool
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		filePaths:   ymlConfig.UList("filePaths"),
		format:      ymlConfig.UBool("format", false),
		formatStyle: ymlConfig.UString("formatStyle", "vim"),
		wrapText:    ymlConfig.UBool("wrapText", true),
	}

	return &settings
}
