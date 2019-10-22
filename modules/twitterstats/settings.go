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

	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string

	screenNames []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		consumerKey:       ymlConfig.UString("consumerKey", os.Getenv("WTF_TWITTER_CONSUMER_KEY")),
		consumerSecret:    ymlConfig.UString("consumerSecret", os.Getenv("WTF_TWITTER_CONSUMER_SECRET")),
		accessToken:       ymlConfig.UString("accessToken", os.Getenv("WTF_TWITTER_ACCESS_TOKEN")),
		accessTokenSecret: ymlConfig.UString("accessTokenSecret", os.Getenv("WTF_TWITTER_ACCESS_TOKEN_SECRET")),

		screenNames: ymlConfig.UList("screenNames"),
	}

	return &settings
}
