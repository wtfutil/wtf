package trello

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "Trello"

type Settings struct {
	common *cfg.Common

	accessToken string
	apiKey      string
	board       string
	list        map[string]string
	username    string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		accessToken: ymlConfig.UString("accessToken", os.Getenv("WTF_TRELLO_ACCESS_TOKEN")),
		apiKey:      ymlConfig.UString("apiKey", os.Getenv("WTF_TRELLO_APP_KEY")),
		board:       ymlConfig.UString("board"),
		username:    ymlConfig.UString("username"),
	}

	settings.list = mapifyList(ymlConfig, globalConfig)

	return &settings
}

func mapifyList(ymlConfig *config.Config, globalConfig *config.Config) map[string]string {
	lists := make(map[string]string)

	// Single list
	list, err := ymlConfig.String("list")
	if err == nil {
		lists[list] = ""
		return lists
	}

	// Array of lists
	listList := ymlConfig.UList("project")
	for _, listName := range listList {
		if list, ok := listName.(string); ok {
			lists[list] = ""
		}
	}

	return lists
}
