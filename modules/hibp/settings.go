package hibp

import (
	"os"
	"time"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable   = false
	defaultTitle       = "HIBP"
	minRefreshInterval = 6 * time.Hour
)

type colors struct {
	ok    string
	pwned string
}

// Settings defines the configuration properties for this module
type Settings struct {
	colors
	*cfg.Common

	accounts []string `help:"A list of the accounts to check the HIBP database for."`
	apiKey   string   `help:"Your HIBP API v3 API key"`
	since    string   `help:"Only check for breaches after this date. Set this if you’ve been breached in the past, have taken steps to mitigate that (changing passwords, cancelling accounts, etc.) and now only want to know about future breaches." values:"A date string in the format 'yyyy-mm-dd', ie. '2019-06-22'" optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_HIBP_TOKEN"))),
		accounts: utils.ToStrs(ymlConfig.UList("accounts")),
		since:    ymlConfig.UString("since", ""),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	settings.colors.ok = ymlConfig.UString("colors.ok", "white")
	settings.colors.pwned = ymlConfig.UString("colors.pwned", "red")

	// HIBP data doesn't need to be reloaded very often so to be gentle on this API we
	// enforce a minimum refresh interval
	if settings.RefreshInterval < minRefreshInterval {
		settings.RefreshInterval = minRefreshInterval
	}

	return settings
}

// HasSince returns TRUE if there's a valid "since" value setting, FALSE if there is not
func (sett *Settings) HasSince() bool {
	if sett.since == "" {
		return false
	}

	_, err := sett.SinceDate()
	return err == nil
}

// SinceDate returns the "since" settings as a proper Time instance
func (sett *Settings) SinceDate() (time.Time, error) {
	dt, err := time.Parse("2006-01-02", sett.since)
	if err != nil {
		return time.Now(), err
	}

	return dt, nil
}
