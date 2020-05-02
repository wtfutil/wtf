package gitlab

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "GitLab"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey   string   `help:"A GitLab personal access token. Requires at least api access."`
	domain   string   `help:"Your GitLab corporate domain."`
	projects []string `help:"A list of key/value pairs each describing a GitLab project to fetch data for." values:"Key: The name of the project. Value: The namespace of the project."`
	username string   `help:"Your GitLab username. Used to figure out which requests require your approval"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_GITLAB_TOKEN"))),
		domain:   ymlConfig.UString("domain"),
		username: ymlConfig.UString("username"),
	}

	cfg.ConfigureSecret(
		globalConfig,
		settings.domain,
		name,
		&settings.username,
		&settings.apiKey,
	)

	settings.projects = cfg.ParseAsMapOrList(ymlConfig, "projects")

	return &settings
}
