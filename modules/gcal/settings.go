package gcal

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "Calendar"

type colors struct {
	day         string
	description string `help:"The default color for calendar event descriptions." values:"Any X11 color name." optional:"true"`
	past        string `help:"The color for calendar events that have passed." values:"Any X11 color name." optional:"true"`
	title       string `help:"The default colour for calendar event titles." values:"Any X11 color name." optional:"true"`

	highlights []interface{} `help:"A list of arrays that define a regular expression pattern and a color. If a calendar event title matches a regular expression, the title will be drawn in that colour. Over-rides the default title colour." values:"An array of a valid regular expression, any X11 color name." optional:"true"`
}

type Settings struct {
	colors
	common *cfg.Common

	conflictIcon          string `help:"The icon displayed beside calendar events that have conflicting times (they intersect or overlap in some way)." values:"Any displayable unicode character." optional:"true"`
	currentIcon           string `help:"The icon displayed beside the current calendar event." values:"Any displayable unicode character." optional:"true"`
	displayResponseStatus bool   `help:"Whether or not to display your response status to the calendar event." values:"true or false" optional:"true"`
	email                 string `help:"The email address associated with your Google account. Necessary for determining 'responseStatus'." values:"A valid email address string."`
	eventCount            int    `help:"The number of calendar events to display." values:"A positive integer, 0..n." optional:"true"`
	multiCalendar         bool   `help:"Whether or not to display your primary calendar or all calendars you have access to." values:"true or false" optional:"true"`
	secretFile        	  string `help:"Your Google client secret JSON file." values:"A string representing a file path to the JSON secret file."`
	showDeclined      	  bool   `help:"Whether or not to display events you’ve declined to attend." values:"true or false" optional:"true"`
	withLocation      	  bool   `help:"Whether or not to show the location of the appointment." values:"true or false"`
	timezone          	  string `help:"The time zone used to display calendar event times." values:"A valid TZ database time zone string" optional:"true"`
	calendarReadLevel 	  string `help:"The calender read level specifies level you want to read events. Default: writer " values:"reader, writer", optional: "true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		conflictIcon:          ymlConfig.UString("conflictIcon", "🚨"),
		currentIcon:           ymlConfig.UString("currentIcon", "🔸"),
		displayResponseStatus: ymlConfig.UBool("displayResponseStatus", true),
		email:                 ymlConfig.UString("email", ""),
		eventCount:            ymlConfig.UInt("eventCount", 10),
		multiCalendar:         ymlConfig.UBool("multiCalendar", false),
		secretFile:            ymlConfig.UString("secretFile", ""),
		showDeclined:          ymlConfig.UBool("showDeclined", false),
		withLocation:          ymlConfig.UBool("withLocation", true),
		timezone:              ymlConfig.UString("timezone", ""),
		calendarReadLevel:     ymlConfig.UString("calendarReadLevel", "writer"),
	}

	settings.colors.day = ymlConfig.UString("colors.day", "forestgreen")
	settings.colors.description = ymlConfig.UString("colors.description", "white")
	settings.colors.past = ymlConfig.UString("colors.past", "gray")
	settings.colors.title = ymlConfig.UString("colors.title", "white")

	return &settings
}
