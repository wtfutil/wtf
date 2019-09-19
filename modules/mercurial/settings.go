package mercurial

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Mercurial"
)

type Settings struct {
	common *cfg.Common

	commitCount  int           `help:"The number of past commits to display." optional:"true"`
	commitFormat string        `help:"The string format for the commit message." optional:"true"`
	repositories []interface{} `help:"Defines which mercurial repositories to watch." values:"A list of zero or more local file paths pointing to valid mercurial repositories."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		commitCount:  ymlConfig.UInt("commitCount", 10),
		commitFormat: ymlConfig.UString("commitFormat", "[forestgreen]{rev}:{phase} [white]{desc|firstline|strip} [grey]{author|person} {date|age}[white]"),
		repositories: ymlConfig.UList("repositories"),
	}

	return &settings
}
