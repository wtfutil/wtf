package twitter

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "twitter"

type Settings struct {
	common *cfg.Common

	bearerToken string
	count       int
	screenNames []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		bearerToken: ymlConfig.UString("bearerToken", os.Getenv("WTF_TWITTER_BEARER_TOKEN")),
		count:       ymlConfig.UInt("count", 5),
		screenNames: ymlConfig.UList("screenName"),
	}

	return &settings
}
