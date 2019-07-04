package gitlab

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "GitLab"

type Settings struct {
	common *cfg.Common

	apiKey   string                 `help:"A GitLab personal access token. Requires at least api access."`
	domain   string                 `help:"Your GitLab corporate domain."`
	projects map[string]interface{} `help:"A list of key/value pairs each describing a GitLab project to fetch data for." values:"Key: The name of the project. Value: The namespace of the project."`
	username string                 `help:"Your GitLab username. Used to figure out which requests require your approval"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", os.Getenv("WTF_GITLAB_TOKEN")),
		domain:   ymlConfig.UString("domain"),
		projects: ymlConfig.UMap("projects"),
		username: ymlConfig.UString("username"),
	}

	return &settings
}
