package clocks

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = false
	defaultTitle     = "Clocks"
)

type colors struct {
	rows struct {
		even string
		odd  string
	}
}

type Settings struct {
	colors
	common *cfg.Common

	dateFormat string                 `help:"The format of the date string for all clocks." values:"Any valid Go date layout which is handled by Time.Format. Defaults to Jan 2."`
	timeFormat string                 `help:"The format of the time string for all clocks." values:"Any valid Go time layout which is handled by Time.Format. Defaults to 15:04 MST."`
	locations  map[string]interface{} `help:"Defines the timezones for the world clocks that you want to display. key is a unique label that will be displayed in the UI. value is a timezone name." values:"Any TZ database timezone."`
	sort       string                 `help:"Defines the display order of the clocks in the widget." values:"'alphabetical' or 'chronological'. 'alphabetical' will sort in acending order by key, 'chronological' will sort in ascending order by date/time."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		dateFormat: ymlConfig.UString("dateFormat", utils.SimpleDateFormat),
		timeFormat: ymlConfig.UString("timeFormat", utils.SimpleTimeFormat),
		locations:  ymlConfig.UMap("locations"),
		sort:       ymlConfig.UString("sort"),
	}

	settings.colors.rows.even = ymlConfig.UString("colors.rows.even", "white")
	settings.colors.rows.odd = ymlConfig.UString("colors.rows.odd", "blue")

	return &settings
}
