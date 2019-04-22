package trello

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "trello"

type Settings struct {
	common *cfg.Common

	accessToken string
	apiKey      string
	board       string
	list        map[string]string
	username    string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		accessToken: localConfig.UString("accessToken", os.Getenv("WTF_TRELLO_ACCESS_TOKEN")),
		apiKey:      localConfig.UString("apiKey", os.Getenv("WTF_TRELLO_APP_KEY")),
		board:       localConfig.UString("board"),
		username:    localConfig.UString("username"),
	}

	settings.list = mapifyList(localConfig)

	return &settings
}

func mapifyList(localConfig *config.Config) map[string]string {
	lists := make(map[string]string)

	// Single list
	list, err := localConfig.String("list")
	if err == nil {
		lists[list] = ""
		return lists
	}

	// Array of lists
	listList := localConfig.UList("project")
	for _, listName := range listList {
		if list, ok := listName.(string); ok {
			lists[list] = ""
		}
	}

	return lists
}
