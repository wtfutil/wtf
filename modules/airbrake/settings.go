package airbrake

import (
	"os"
	"strconv"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Airbrake"
)

type Settings struct {
	*cfg.Common

	projectID int    `help:"The id of your Airbrake project."`
	authToken string `help:"The token that allows accessing Airbrake API"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle,
			defaultFocusable, ymlConfig, globalConfig),
		projectID: ymlConfig.UInt("projectID", getProjectID()),
		authToken: ymlConfig.UString("authToken", os.Getenv("AIRBRAKE_USER_KEY")),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.authToken).Load()

	return &settings
}

func getProjectID() int {
	projectID, err := strconv.ParseInt(os.Getenv("AIRBRAKE_PROJECT_ID"), 10, 32)
	if err != nil {
		return 0
	}

	return int(projectID)
}
