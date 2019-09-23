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

	labelColor  string
	apiToken    string `help:"Your Azure DevOps Access Token."`
	orgURL      string `help:"Your Azure DevOps organization URL."`
	projectName string
	maxRows     int
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocus, ymlConfig, globalConfig),

		labelColor:  ymlConfig.UString("labelColor", "white"),
		apiToken:    ymlConfig.UString("apiToken", os.Getenv("WTF_AZURE DEVOPS_API_TOKEN")),
		orgURL:      ymlConfig.UString("orgURL", os.Getenv("WTF_AZURE_DEVOPS_ORG_URL")),
		projectName: ymlConfig.UString("projectName", os.Getenv("WTF_AZURE_DEVOPS_PROJECT_NAME")),
		maxRows:     ymlConfig.UInt("maxRows", 3),
	}

	return &settings
}
