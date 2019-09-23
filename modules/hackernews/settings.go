package hackernews

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "HackerNews"
)

type Settings struct {
	common *cfg.Common

	numberOfStories int    `help:"Defines number of stories to be displayed. Default is 10" optional:"true"`
	storyType       string `help:"Category of story to see" values:"new, top, job, ask" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		numberOfStories: ymlConfig.UInt("numberOfStories", 10),
		storyType:       ymlConfig.UString("storyType", "top"),
	}

	return &settings
}
