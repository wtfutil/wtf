package datadog

import (
	"github.com/wtfutil/wtf/utils"
	datadog "github.com/zorkian/go-datadog-api"
)

// Monitors returns a list of newrelic monitors
func (widget *Widget) Monitors() ([]datadog.Monitor, error) {
	client := datadog.NewClient(
		widget.settings.apiKey,
		widget.settings.applicationKey,
	)

	tags := utils.ToStrs(widget.settings.tags)

	monitors, err := client.GetMonitorsByTags(tags)
	if err != nil {
		return nil, err
	}

	return monitors, nil
}
