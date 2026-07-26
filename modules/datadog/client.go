package datadog

import (
	"context"
	"strings"

	ddapi "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/wtfutil/wtf/utils"
)

// Monitors returns a list of Datadog monitors
func (widget *Widget) Monitors() ([]datadogV1.Monitor, error) {
	ctx := context.WithValue(context.Background(), ddapi.ContextAPIKeys, map[string]ddapi.APIKey{
		"apiKeyAuth": {Key: widget.settings.apiKey},
		"appKeyAuth": {Key: widget.settings.applicationKey},
	})

	config := ddapi.NewConfiguration()
	client := ddapi.NewAPIClient(config)

	tags := utils.ToStrs(widget.settings.tags)

	optionalParams := datadogV1.NewListMonitorsOptionalParameters()
	if len(tags) > 0 {
		tagFilter := strings.Join(tags, ",")
		optionalParams = optionalParams.WithMonitorTags(tagFilter)
	}

	monitorsApi := datadogV1.NewMonitorsApi(client)
	monitors, _, err := monitorsApi.ListMonitors(ctx, *optionalParams)
	if err != nil {
		return nil, err
	}

	return monitors, nil
}
