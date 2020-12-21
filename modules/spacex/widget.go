package spacex

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	view.TextWidget
	settings *Settings
	err      error
}

func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),
		settings:   settings,
	}
	return widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}
	widget.Redraw(widget.content)
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	var title = "Next SpaceX 🚀"
	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}

	launch, err := NextLaunch()
	var str string
	if err != nil {
		handleError(widget, err)
	} else {

		str = fmt.Sprintf("[%s]Mission[white]\n", widget.settings.Colors.Subheading)
		str += fmt.Sprintf("%s: %s\n", "Name", launch.MissionName)
		str += fmt.Sprintf("%s: %s\n", "Date", wtf.UnixTime(launch.LaunchDate).Format(time.RFC822))
		str += fmt.Sprintf("%s: %s\n", "Site", launch.LaunchSite.Name)
		str += "\n"

		str += fmt.Sprintf("[%s]Links[white]\n", widget.settings.Colors.Subheading)
		str += fmt.Sprintf("%s: %s\n", "YouTube", launch.Links.YouTubeLink)
		str += fmt.Sprintf("%s: %s\n", "Reddit", launch.Links.RedditLink)

		if widget.CommonSettings().Height >= 2 {
			str += "\n"
			str += fmt.Sprintf("[%s]Details[white]\n", widget.settings.Colors.Subheading)
			str += fmt.Sprintf("%s: %s\n", "RocketName", launch.Rocket.Name)
			str += fmt.Sprintf("%s: %s\n", "Details", launch.Details)
		}
	}
	return title, str, true
}

func handleError(widget *Widget, err error) {
	widget.err = err
}
