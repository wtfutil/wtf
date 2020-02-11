package cdsstatus

import (
	"fmt"
	"strings"

	"github.com/ovh/cds/sdk"
)

func (widget *Widget) display() {
	widget.TextWidget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	if len(widget.View.GetHighlights()) > 0 {
		widget.View.ScrollToHighlight()
	} else {
		widget.View.ScrollToBeginning()
	}

	widget.Items = make([]sdk.MonitoringStatusLine, 0)
	str := widget.displayStatus()
	title := widget.CommonSettings().Title
	return title, str, false
}

func (widget *Widget) displayStatus() string {
	status, err := widget.client.MonStatus()

	if err != nil || len(status.Lines) == 0 {
		return fmt.Sprintf(" [red]Error: %v[white]\n", err.Error())
	}

	widget.SetItemCount(len(status.Lines))

	var (
		global     []string
		globalWarn []string
		globalRed  []string
		ok         []string
		warn       []string
		red        []string
	)

	for _, line := range status.Lines {
		if line.Status == sdk.MonitoringStatusWarn && strings.Contains(line.Component, "Global") {
			globalWarn = append(globalWarn, line.String())
		} else if line.Status != sdk.MonitoringStatusOK && strings.Contains(line.Component, "Global") {
			globalRed = append(globalRed, line.String())
		} else if strings.Contains(line.Component, "Global") {
			global = append(global, line.String())
		} else if line.Status == sdk.MonitoringStatusWarn {
			warn = append(warn, line.String())
		} else if line.Status == sdk.MonitoringStatusOK {
			ok = append(ok, line.String())
		} else {
			red = append(red, line.String())
		}
	}
	var idx int
	var content string
	for _, v := range globalRed {
		content += fmt.Sprintf("[grey][\"%d\"][red]%s\n", idx, v)
		idx++
	}
	for _, v := range globalWarn {
		content += fmt.Sprintf("[grey][\"%d\"][yellow]%s\n", idx, v)
		idx++
	}
	for _, v := range global {
		content += fmt.Sprintf("[grey][\"%d\"][grey]%s\n", idx, v)
		idx++
	}
	for _, v := range red {
		content += fmt.Sprintf("[grey][\"%d\"][red]%s\n", idx, v)
		idx++
	}
	for _, v := range warn {
		content += fmt.Sprintf("[grey][\"%d\"][yellow]%s\n", idx, v)
		idx++
	}
	for _, v := range ok {
		content += fmt.Sprintf("[grey][\"%d\"][grey]%s\n", idx, v)
		idx++
	}
	return content
}

func getStatusColor(status string) string {
	switch status {
	case sdk.StatusSuccess:
		return "green"
	case sdk.StatusBuilding, sdk.StatusWaiting:
		return "blue"
	case sdk.StatusFail:
		return "red"
	case sdk.StatusStopped:
		return "red"
	case sdk.StatusSkipped:
		return "grey"
	case sdk.StatusDisabled:
		return "grey"
	}
	return "red"
}

func pad(t string, size int) string {
	if len(t) > size {
		return t[0:size-3] + "..."
	}
	return t + strings.Repeat(" ", size-len(t))
}

func getVarsInPbj(key string, ps []sdk.Parameter) string {
	for _, p := range ps {
		if p.Name == key {
			return p.Value
		}
	}
	return ""
}
