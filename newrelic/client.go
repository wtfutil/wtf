package newrelic

import (
	"os"

	nr "github.com/yfronto/newrelic"
)

func Application() (*nr.Application, error) {
	client := nr.NewClient(os.Getenv("WTF_NEW_RELIC_API_KEY"))

	application, err := client.GetApplication(Config.UInt("wtf.newrelic.applicationId"))
	if err != nil {
		return nil, err
	}

	return application, nil
}

func Deployments() ([]nr.ApplicationDeployment, error) {
	client := nr.NewClient(os.Getenv("WTF_NEW_RELIC_API_KEY"))

	opts := &nr.ApplicationDeploymentOptions{Page: 1}
	deployments, err := client.GetApplicationDeployments(Config.UInt("wtf.newrelic.applicationId"), opts)
	if err != nil {
		return nil, err
	}

	return deployments, nil
}
