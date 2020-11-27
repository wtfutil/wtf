package todo_plus

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultTitle     = "Todo"
	defaultFocusable = true
)

type Settings struct {
	*cfg.Common

	backendType     string
	backendSettings *config.Config
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	backend, _ := ymlConfig.Get("backendSettings")

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		backendType:     ymlConfig.UString("backendType"),
		backendSettings: backend,
	}

	return &settings
}

func FromTodoist(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	apiKey := ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TODOIST_TOKEN")))
	cfg.ModuleSecret(name, globalConfig, &apiKey).Load()
	projects := ymlConfig.UList("projects")
	backend, _ := config.ParseYaml("apiKey: " + apiKey)
	_ = backend.Set(".projects", projects)

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		backendType:     "todoist",
		backendSettings: backend,
	}

	return &settings
}

func FromTrello(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	accessToken := ymlConfig.UString("accessToken", ymlConfig.UString("apikey", os.Getenv("WTF_TRELLO_ACCESS_TOKEN")))
	apiKey := ymlConfig.UString("apiKey", os.Getenv("WTF_TRELLO_API_KEY"))
	cfg.ModuleSecret(name, globalConfig, &apiKey).Load()
	board := ymlConfig.UString("board")
	username := ymlConfig.UString("username")
	var lists []interface{}
	list, err := ymlConfig.String("list")
	if err == nil {
		lists = append(lists, list)
	} else {
		lists = ymlConfig.UList("list")
	}
	backend, _ := config.ParseYaml("apiKey: " + apiKey)
	_ = backend.Set(".accessToken", accessToken)
	_ = backend.Set(".board", board)
	_ = backend.Set(".username", username)
	_ = backend.Set(".lists", lists)

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		backendType:     "trello",
		backendSettings: backend,
	}

	return &settings
}
