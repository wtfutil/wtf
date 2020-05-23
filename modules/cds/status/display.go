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
		switch {
		case line.Status == sdk.MonitoringStatusWarn && strings.Contains(line.Component, "Global"):
			globalWarn = append(globalWarn, line.String())
		case line.Status != sdk.MonitoringStatusOK && strings.Contains(line.Component, "Global"):
			globalRed = append(globalRed, line.String())
		case strings.Contains(line.Component, "Global"):
			global = append(global, line.String())
		case line.Status == sdk.MonitoringStatusWarn:
			warn = append(warn, line.String())
		case line.Status == sdk.MonitoringStatusOK:
			ok = append(ok, line.String())
		default:
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
