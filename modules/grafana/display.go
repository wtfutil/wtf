package grafana

import "fmt"

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title

	var out string
	if widget.Err != nil {
		return title, widget.Err.Error(), false
	} else {
		for idx, alert := range widget.Alerts {
			out += fmt.Sprintf(` ["%d"][%s]%s - %s[""]`,
				idx,
				stateColor(alert.State),
				stateToEmoji(alert.State),
				alert.Name,
			)
			out += "\n"
		}
	}

	return title, out, false
}

func stateColor(state AlertState) string {
	switch state {
	case Ok:
		return "green"
	case Paused:
		return "yellow"
	case Alerting:
		return "red"
	case Pending:
		return "orange"
	case NoData:
		return "yellow"
	default:
		return "white"
	}
}

func stateToEmoji(state AlertState) string {
	switch state {
	case Ok:
		return "✔"
	case Paused:
		return "⏸"
	case Alerting:
		return "✘"
	case Pending:
		return "?"
	case NoData:
		return "?"
	}
	return ""
}
