package datadog

import (
	"os"

	"github.com/senorprogrammer/wtf/wtf"
	datadog "github.com/zorkian/go-datadog-api"
)

// Monitors returns a list of newrelic monitors
func Monitors() ([]datadog.Monitor, error) {
	client := datadog.NewClient(apiKey(), applicationKey())

	monitors, err := client.GetMonitorsByTags(wtf.ToStrs(wtf.Config.UList("wtf.mods.datadog.monitors.tags")))
	if err != nil {
		return nil, err
	}

	return monitors, nil
}

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.datadog.apiKey",
		os.Getenv("WTF_DATADOG_API_KEY"),
	)
}

func applicationKey() string {
	return wtf.Config.UString(
		"wtf.mods.datadog.applicationKey",
		os.Getenv("WTF_DATADOG_APPLICATION_KEY"),
	)
}
