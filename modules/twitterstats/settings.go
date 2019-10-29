package twitterstats

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Twitter Stats"
)

type Settings struct {
	common *cfg.Common

	bearerToken    string
	consumerKey    string
	consumerSecret string
	screenNames    []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		bearerToken:    ymlConfig.UString("bearerToken", os.Getenv("WTF_TWITTER_BEARER_TOKEN")),
		consumerKey:    ymlConfig.UString("consumerKey", os.Getenv("WTF_TWITTER_CONSUMER_KEY")),
		consumerSecret: ymlConfig.UString("consumerSecret", os.Getenv("WTF_TWITTER_CONSUMER_SECRET")),

		screenNames: ymlConfig.UList("screenNames"),
	}

	return &settings
}
