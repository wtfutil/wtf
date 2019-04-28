package newrelic

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "newrelic"

type Settings struct {
	common *cfg.Common

	apiKey        string
	applicationID int
	deployCount   int
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:        ymlConfig.UString("apiKey", os.Getenv("WTF_NEW_RELIC_API_KEY")),
		applicationID: ymlConfig.UInt("applicationID"),
		deployCount:   ymlConfig.UInt("deployCount", 5),
	}

	return &settings
}
