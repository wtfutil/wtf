package toggl

import (
	"fmt"
	"sort"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	*Client

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),
		Client:     NewClient(settings.apiKey),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	me, err := widget.Client.me()

	title := fmt.Sprintf("%s - Time Entries", widget.CommonSettings().Title)
	var str, color, durationStr, entryFormat string

	wrap := false
	if err != nil {
		wrap = true
		str = err.Error()
	} else {

		sort.Slice(me.Data.Time_entries, func(i, j int) bool {
			idA := me.Data.Time_entries[i].Id
			idB := me.Data.Time_entries[j].Id
			return idA > idB
		})

		for idx, entries := range me.Data.Time_entries {
			if idx >= widget.settings.numberOfEntries {
				break
			}

			//	start and stop into times
			startTime, _ := time.Parse(
				time.RFC3339,
				entries.Start)

			stopTime, _ := time.Parse(
				time.RFC3339,
				entries.Stop)

			if entries.Duration < 0 {
				color = widget.settings.runningEntryColor
				secs, _ := time.ParseDuration(fmt.Sprintf("%d%s", time.Now().Unix()+int64(entries.Duration), "s"))
				durationStr = secs.String()
				entryFormat = "[%s] %s[dark%[1]s] from %[3]s UTC [%[5]s](Duration: %[6]s)\n"
			} else {
				color = widget.settings.completedEntryColor
				secs, _ := time.ParseDuration(fmt.Sprintf("%d%s", entries.Duration, "s"))
				durationStr = secs.String()
				entryFormat = "[%s] %s[dark%[1]s] from %[3]s to %[4]s UTC [%[5]s](Duration: %[6]s)\n"
			}

			str += fmt.Sprintf(
				entryFormat,
				color,
				entries.Description,
				fmt.Sprintf(startTime.Format(time.Stamp)),
				fmt.Sprintf(stopTime.Format(time.Stamp)),
				widget.settings.durationColor,
				durationStr,
			)
		}
	}

	return title, str, wrap
}
