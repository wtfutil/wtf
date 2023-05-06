package digitalclock

import "strings"

func mergeLines(outString []string) string {
	return strings.Join(outString, "\n")
}

func renderWidget(widgetSettings Settings) string {
	outputStrings := []string{}

	clockString, needBorder := renderClock(widgetSettings)
	if needBorder {
		outputStrings = append(outputStrings, mergeLines([]string{"", clockString, ""}))
	} else {
		outputStrings = append(outputStrings, clockString)
	}

	if widgetSettings.withDate {
		outputStrings = append(outputStrings, getDate(widgetSettings.dateFormat, widgetSettings.withDatePrefix))
	}

	if widgetSettings.withUTC {
		outputStrings = append(outputStrings, getUTC())
	}

	if widgetSettings.withEpoch {
		outputStrings = append(outputStrings, getEpoch())
	}

	return mergeLines(outputStrings)
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		title := widget.CommonSettings().Title
		if widget.settings.dateTitle {
			title = getDate(widget.settings.dateFormat, false)
		}
		return title, renderWidget(*widget.settings), false
	})
}
