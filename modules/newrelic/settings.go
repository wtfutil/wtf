package newrelic

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "NewRelic"
)

type Settings struct {
	common *cfg.Common

	apiKey         string        `help:"Your New Relic API token."`
	deployCount    int           `help:"The number of past deploys to display on screen." optional:"true"`
	applicationIDs []interface{} `help:"The integer ID of the New Relic application you wish to report on."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:         ymlConfig.UString("apiKey", os.Getenv("WTF_NEW_RELIC_API_KEY")),
		deployCount:    ymlConfig.UInt("deployCount", 5),
		applicationIDs: ymlConfig.UList("applicationIDs"),
	}

	return &settings
}
