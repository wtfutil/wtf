package azuredevops

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocus = false
	defaultTitle = "azuredevops"
)

// Settings defines the configuration options for this module
type Settings struct {
	common *cfg.Common

	apiToken    string `help:"Your Azure DevOps Access Token."`
	labelColor  string
	maxRows     int
	orgURL      string `help:"Your Azure DevOps organization URL."`
	projectName string
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocus, ymlConfig, globalConfig),

		apiToken:    ymlConfig.UString("apiToken", os.Getenv("WTF_AZURE_DEVOPS_API_TOKEN")),
		labelColor:  ymlConfig.UString("labelColor", "white"),
		maxRows:     ymlConfig.UInt("maxRows", 3),
		orgURL:      ymlConfig.UString("orgURL", os.Getenv("WTF_AZURE_DEVOPS_ORG_URL")),
		projectName: ymlConfig.UString("projectName", os.Getenv("WTF_AZURE_DEVOPS_PROJECT_NAME")),
	}

	cfg.ConfigureSecret(
		globalConfig,
		settings.orgURL,
		"",
		&settings.projectName,
		&settings.apiToken,
	)

	return &settings
}
