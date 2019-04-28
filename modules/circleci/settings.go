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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey: ymlConfig.UString("apiKey", os.Getenv("WTF_CIRCLE_API_KEY")),
	}

	return &settings
}
