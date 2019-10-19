package football

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "football"
)

type Settings struct {
	common *cfg.Common
	apiKey string `help:"Your Football-data API token."`
	league string `help:"Name of the competition. For example PL"`
	team   string `help:"Teams to follow in mentioned league"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		apiKey: ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_FOOTBALL_API_KEY"))),
		league: ymlConfig.UString("league", ymlConfig.UString("league", os.Getenv("WTF_FOOTBALL_LEAGUE"))),
		team:   ymlConfig.UString("teams", ymlConfig.UString("teams", os.Getenv("WTF_FOOTBALL_TEAM"))),
	}
	return &settings
}
