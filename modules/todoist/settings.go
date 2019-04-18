package todoist

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "todoist"

type Settings struct {
	common *cfg.Common

	apiKey   string
	projects []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey:   localConfig.UString("apiKey", os.Getenv("WTF_TODOIST_TOKEN")),
		projects: localConfig.UList("projects"),
	}

	return &settings
}
