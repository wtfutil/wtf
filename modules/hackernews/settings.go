package hackernews

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "hackernews"

type Settings struct {
	common *cfg.Common

	numberOfStories int
	storyType       string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		numberOfStories: localConfig.UInt("numberOfStories", 10),
		storyType:       localConfig.UString("storyType", "top"),
	}

	return &settings
}
