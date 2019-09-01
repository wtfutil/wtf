package todo_plus

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "Todo"

type Settings struct {
	common *cfg.Common

	backendType     string
	backendSettings *config.Config
	projects        []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	backend, _ := ymlConfig.Get("backendSettings")

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		backendType:     ymlConfig.UString("backendType"),
		backendSettings: backend,
		projects:        ymlConfig.UList("projects"),
	}

	return &settings
}

func FromTodoist(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	apiKey := ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TODOIST_TOKEN")))
	backend, _ := config.ParseYaml("apiKey: " + apiKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		backendType:     "todoist",
		backendSettings: backend,
		projects:        ymlConfig.UList("projects"),
	}

	return &settings
}
