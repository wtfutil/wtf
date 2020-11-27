package gcal

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Calendar"
)

type colors struct {
	day         string
	description string `help:"The default color for calendar event descriptions." values:"Any X11 color name." optional:"true"`
	eventTime   string `help:"The default color for calendar event times." values:"Any X11 color name." optional:"true"`
	past        string `help:"The color for calendar events that have passed." values:"Any X11 color name." optional:"true"`
	title       string `help:"The default colour for calendar event titles." values:"Any X11 color name." optional:"true"`

	highlights []interface{} `help:"A list of arrays that define a regular expression pattern and a color. If a calendar event title matches a regular expression, the title will be drawn in that colour. Over-rides the default title colour." values:"An array of a valid regular expression, any X11 color name." optional:"true"`
}

// Settings defines the configuration options for this module
type Settings struct {
	colors
	*cfg.Common

	conflictIcon          string `help:"The icon displayed beside calendar events that have conflicting times (they intersect or overlap in some way)." values:"Any displayable unicode character." optional:"true"`
	currentIcon           string `help:"The icon displayed beside the current calendar event." values:"Any displayable unicode character." optional:"true"`
	displayResponseStatus bool   `help:"Whether or not to display your response status to the calendar event." values:"true or false" optional:"true"`
	email                 string `help:"The email address associated with your Google account. Necessary for determining 'responseStatus'." values:"A valid email address string."`
	eventCount            int    `help:"The number of calendar events to display." values:"A positive integer, 0..n." optional:"true"`
	hourFormat            string `help:"The format of the clock." values:"12 or 24"`
	multiCalendar         bool   `help:"Whether or not to display your primary calendar or all calendars you have access to." values:"true or false" optional:"true"`
	secretFile            string `help:"Your Google client secret JSON file." values:"A string representing a file path to the JSON secret file."`
	showAllDay            bool   `help:"Whether or not to display all-day events" values:"true or false" optional:"true" default:"true"`
	showDeclined          bool   `help:"Whether or not to display events you’ve declined to attend." values:"true or false" optional:"true"`
	showEndTime           bool   `help:"Display the end time of events, in addition to start time." values:"true or false" optional:"true" default:"false"`
	withLocation          bool   `help:"Whether or not to show the location of the appointment." values:"true or false"`
	timezone              string `help:"The time zone used to display calendar event times." values:"A valid TZ database time zone string" optional:"true"`
	calendarReadLevel     string `help:"The calender read level specifies level you want to read events. Default: writer " values:"reader, writer" optional:"true"`
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		conflictIcon:          ymlConfig.UString("conflictIcon", "🚨"),
		currentIcon:           ymlConfig.UString("currentIcon", "🔸"),
		displayResponseStatus: ymlConfig.UBool("displayResponseStatus", true),
		email:                 ymlConfig.UString("email", ""),
		eventCount:            ymlConfig.UInt("eventCount", 10),
		hourFormat:            ymlConfig.UString("hourFormat", "24"),
		multiCalendar:         ymlConfig.UBool("multiCalendar", false),
		secretFile:            ymlConfig.UString("secretFile", ""),
		showAllDay:            ymlConfig.UBool("showAllDay", true),
		showEndTime:           ymlConfig.UBool("showEndTime", false),
		showDeclined:          ymlConfig.UBool("showDeclined", false),
		withLocation:          ymlConfig.UBool("withLocation", true),
		timezone:              ymlConfig.UString("timezone", ""),
		calendarReadLevel:     ymlConfig.UString("calendarReadLevel", "writer"),
	}

	settings.colors.day = ymlConfig.UString("colors.day", settings.Colors.Subheading)
	settings.colors.description = ymlConfig.UString("colors.description", "white")

	// settings.colors.eventTime is a new feature introduced via issue #638. Prior to this, the color of the event
	// time was (unintentionally) customized via settings.colors.description. To maintain backwards compatibility
	// for users who might be already using this to set the color of the event time, we try to determine the default
	// from settings.colors.description. If it is not set, then the default value of "white" is used.  Finally, if a
	// user sets a value for colors.eventTime, it overrides the defaults.
	//
	// PS: We should have a deprecation plan for supporting this backwards compatibility feature.
	settings.colors.eventTime = ymlConfig.UString("colors.eventTime", settings.colors.description)

	settings.colors.highlights = ymlConfig.UList("colors.highlights")
	settings.colors.past = ymlConfig.UString("colors.past", "gray")
	settings.colors.title = ymlConfig.UString("colors.title", "white")

	settings.SetDocumentationPath("google/gcal")

	return &settings
}
