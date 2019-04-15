package hackernews

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	numberOfStories int
	storyType       string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.hackernews")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		numberOfStories: localConfig.UInt("numberOfStories", 10),
		storyType:       localConfig.UString("storyType", "top"),
	}

	return &settings
}
