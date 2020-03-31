package football

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "football"
)

type Settings struct {
	common        *cfg.Common
	apiKey        string `help:"Your Football-data API token."`
	league        string `help:"Name of the competition. For example PL"`
	favTeam       string `help:"Teams to follow in mentioned league"`
	matchesFrom   int    `help:"Matches till Today (Today - Number of days), Default: 2"`
	matchesTo     int    `help:"Matches from Today (Today + Number of days), Default: 5"`
	standingCount int    `help:"Top N number of teams in standings, Default: 5"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common:        cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		apiKey:        ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_FOOTBALL_API_KEY"))),
		league:        ymlConfig.UString("league", ymlConfig.UString("league", os.Getenv("WTF_FOOTBALL_LEAGUE"))),
		favTeam:       ymlConfig.UString("favTeam", ymlConfig.UString("favTeam", os.Getenv("WTF_FOOTBALL_TEAM"))),
		matchesFrom:   ymlConfig.UInt("matchesFrom", 5),
		matchesTo:     ymlConfig.UInt("matchesTo", 5),
		standingCount: ymlConfig.UInt("standingCount", 5),
	}

	cfg.ConfigureSecret(
		globalConfig,
		"",
		name,
		nil,
		&settings.apiKey,
	)

	return &settings
}
