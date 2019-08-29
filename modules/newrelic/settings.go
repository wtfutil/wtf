package newrelic

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "NewRelic"

type Settings struct {
	common *cfg.Common

	apiKey        string `help:"Your New Relic API token."`
	applicationID int    `help:"The integer ID of the New Relic application you wish to report on."`
	deployCount   int    `help:"The number of past deploys to display on screen." optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:        ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_NEW_RELIC_API_KEY"))),
		applicationID: ymlConfig.UInt("applicationID"),
		deployCount:   ymlConfig.UInt("deployCount", 5),
	}

	return &settings
}
