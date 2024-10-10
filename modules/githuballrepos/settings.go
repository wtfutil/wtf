package githuballrepos

import (
	"os/user"
	"path/filepath"
	"strings"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	APIKey        string   `help:"Your GitHub API token."`
	Organizations []string `help:"List of GitHub organizations and user accounts to monitor."`
	Username      string   `help:"Your GitHub username."`
	CachePath     string   `help:"The path to the cache file for this module."`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		common: cfg.NewCommonSettingsFromModule(name, "githuballrepos", true, ymlConfig, globalConfig),

		APIKey:    ymlConfig.UString("apiKey", ""),
		Username:  ymlConfig.UString("username", ""),
		CachePath: expandHomeDir(ymlConfig.UString("cache", "~/.config/wtf/.cache/githuballrepos/")),
	}

	// Convert []interface{} to []string for Organizations
	orgInterfaces := ymlConfig.UList("organizations")
	settings.Organizations = make([]string, len(orgInterfaces))
	for i, org := range orgInterfaces {
		settings.Organizations[i] = org.(string)
	}

	return settings
}

// expandHomeDir replaces the `~` in the path with the user's home directory
func expandHomeDir(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			return path
		}
		return filepath.Join(usr.HomeDir, path[2:])
	}
	return path
}
