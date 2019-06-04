package newrelic

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "NewRelic"

type Settings struct {
	common *cfg.Common

	apiKey         string
	deployCount    int
	applicationIDs []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:         ymlConfig.UString("apiKey", os.Getenv("WTF_NEW_RELIC_API_KEY")),
		deployCount:    ymlConfig.UInt("deployCount", 5),
		applicationIDs: ymlConfig.UList("applicationIDs"),
	}

	return &settings
}
