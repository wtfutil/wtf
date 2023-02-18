package tmuxinator

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Tmuxinator Projects"
)

type Settings struct {
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	common := cfg.NewCommonSettingsFromModule(
		name,
		defaultTitle,
		defaultFocusable,
		ymlConfig,
		globalConfig,
	)

	settings := Settings{common: common}

	return &settings
}
