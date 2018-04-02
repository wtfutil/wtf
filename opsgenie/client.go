package opsgenie

import (
	"os"

	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	sch "github.com/opsgenie/opsgenie-go-sdk/schedule"
)

func Fetch() string {
	apiKey := os.Getenv("WTF_OPS_GENIE_API_KEY")

	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(apiKey)

	scheduler, err := cli.Schedule()
	if err != nil {
		panic(err)
	}

	request := sch.ListSchedulesRequest{}
	response, err := scheduler.List(request)
	if err != nil {
		panic(err)
	}

	var str string
	for _, schedule := range response.Schedules {
		str = str + schedule.Name + "\n"
	}

	return ""
}
