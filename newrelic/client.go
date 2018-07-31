package newrelic

import (
	"os"

	"github.com/senorprogrammer/wtf/wtf"
	nr "github.com/yfronto/newrelic"
)

func Application() (*nr.Application, error) {
	client := nr.NewClient(apiKey())

	application, err := client.GetApplication(wtf.Config.UInt("wtf.mods.newrelic.applicationId"))
	if err != nil {
		return nil, err
	}

	return application, nil
}

func Deployments() ([]nr.ApplicationDeployment, error) {
	client := nr.NewClient(apiKey())

	opts := &nr.ApplicationDeploymentOptions{Page: 1}
	deployments, err := client.GetApplicationDeployments(wtf.Config.UInt("wtf.mods.newrelic.applicationId"), opts)
	if err != nil {
		return nil, err
	}

	return deployments, nil
}

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.newrelic.apiKey",
		os.Getenv("WTF_NEW_RELIC_API_KEY"),
	)
}
