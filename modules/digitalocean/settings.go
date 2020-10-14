package digitalocean

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

const (
	defaultFocusable = true
	defaultTitle     = "DigitalOcean"
)

// defaultColumns defines the default set of columns to display in the widget
// This can be over-ridden in the cofig by explicitly defining a set of columns
var defaultColumns = []interface{}{
	"Name",
	"Status",
	"Region.Slug",
}

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	apiKey     string   `help:"Your DigitalOcean API key."`
	columns    []string `help:"A list of the droplet properties to display."`
	dateFormat string   `help:"The format to display dates and times in."`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:     ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_DIGITALOCEAN_API_KEY"))),
		columns:    utils.ToStrs(ymlConfig.UList("columns", defaultColumns)),
		dateFormat: ymlConfig.UString("dateFormat", wtf.DateFormat),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
