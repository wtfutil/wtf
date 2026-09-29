package docker

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "docker"
)

// Settings defines the configuration options for this module
type Settings struct {
	*cfg.Common

	labelColor  string
	pidFilePath string
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common:     cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		labelColor: ymlConfig.UString("labelColor", "white"),

		// Path to dockerd's pid file, read instead of probing the API so a
		// socket-activated daemon isn't woken up. "auto" derives it; "" disables.
		pidFilePath: ymlConfig.UString("pidFilePath", ""),
	}

	return &settings
}

func (widget *Widget) ConfigText() string {
	return utils.HelpFromInterface(Settings{})
}
