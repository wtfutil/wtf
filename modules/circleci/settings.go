package circleci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "circleci"

type Settings struct {
	common *cfg.Common

	apiKey string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey: localConfig.UString("apiKey", os.Getenv("WTF_CIRCLE_API_KEY")),
	}

	return &settings
}
