package git

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Git"
)

type Settings struct {
	common *cfg.Common

	commitCount  int           `help:"The number of past commits to display." values:"A positive integer, 0..n." optional:"true"`
	commitFormat string        `help:"The string format for the commit message." optional:"true"`
	dateFormat   string        `help:"The string format for the date/time in the commit message." optional:"true"`
	repositories []interface{} `help:"Defines which git repositories to watch." values:"A list of zero or more local file paths pointing to valid git repositories."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		commitCount:  ymlConfig.UInt("commitCount", 10),
		commitFormat: ymlConfig.UString("commitFormat", "[forestgreen]%h [white]%s [grey]%an on %cd[white]"),
		dateFormat:   ymlConfig.UString("dateFormat", "%b %d, %Y"),
		repositories: ymlConfig.UList("repositories"),
	}

	return &settings
}

func (widget *Widget) ConfigText() string {
	return utils.HelpFromInterface(Settings{})
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}
