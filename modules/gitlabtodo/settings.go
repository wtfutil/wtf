package gitlabtodo

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "GitLab Todos"
)

type Settings struct {
	common *cfg.Common

	numberOfTodos int    `help:"Defines number of stories to be displayed. Default is 10" optional:"true"`
	apiKey        string `help:"A GitLab personal access token. Requires at least api access."`
	domain        string `help:"Your GitLab corporate domain."`
	showProject   bool   `help:"Determines whether or not to show the project a given todo is for."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		numberOfTodos: ymlConfig.UInt("numberOfTodos", 10),
		apiKey:        ymlConfig.UString("apiKey", os.Getenv("WTF_GITLAB_TOKEN")),
		domain:        ymlConfig.UString("domain"),
		showProject:   ymlConfig.UBool("showProject", true),
	}

	cfg.ConfigureSecret(
		globalConfig,
		settings.domain,
		name,
		nil,
		&settings.apiKey,
	)

	return &settings
}
