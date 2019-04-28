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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		numberOfStories: ymlConfig.UInt("numberOfStories", 10),
		storyType:       ymlConfig.UString("storyType", "top"),
	}

	return &settings
}
