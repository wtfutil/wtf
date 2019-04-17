package twitter

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	bearerToken string
	count       int
	screenNames []interface{}
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.twitter")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		bearerToken: localConfig.UString("bearerToken", os.Getenv("WTF_TWITTER_BEARER_TOKEN")),
		count:       localConfig.UInt("count", 5),
		screenNames: localConfig.UList("screenName"),
	}

	return &settings
}
