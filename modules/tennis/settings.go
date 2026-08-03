package tennis

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Tennis"

	defaultStatus     = "live"
	defaultMatchLimit = 10
)

// Settings defines the configuration options for this module
type Settings struct {
	*cfg.Common

	apiKey     string `help:"Your Live Tennis API key. A free key (1,000 requests/day) is available at https://livetennisapi.com/subscribe/free." values:"A valid Live Tennis API key"`
	baseURL    string `help:"The base URL of the Live Tennis API." values:"A URL" optional:"true" default:"https://api.livetennisapi.com/api/public/v1"`
	tour       string `help:"Restrict matches to a single tour." values:"atp, wta, or empty for all tours" optional:"true"`
	status     string `help:"Which matches to display." values:"live, upcoming, completed" optional:"true" default:"live"`
	matchLimit int    `help:"The maximum number of matches to display." values:"A positive integer" optional:"true" default:"10"`
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:     ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TENNIS_API_KEY"))),
		baseURL:    ymlConfig.UString("baseURL", defaultBaseURL),
		tour:       ymlConfig.UString("tour", ""),
		status:     normalizeStatus(ymlConfig.UString("status", defaultStatus)),
		matchLimit: ymlConfig.UInt("matchLimit", defaultMatchLimit),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	settings.SetDocumentationPath("sports/tennis")

	return &settings
}

// normalizeStatus clamps the configured status filter to one the API accepts,
// falling back to the default rather than sending a bad request.
func normalizeStatus(status string) string {
	for _, valid := range statusCycle {
		if status == valid {
			return status
		}
	}

	return defaultStatus
}

func (widget *Widget) ConfigText() string {
	return utils.HelpFromInterface(Settings{})
}
