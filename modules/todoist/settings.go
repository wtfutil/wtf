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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", os.Getenv("WTF_TODOIST_TOKEN")),
		projects: ymlConfig.UList("projects"),
	}

	return &settings
}
