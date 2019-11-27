package newrelic

import (
	nr "github.com/wtfutil/wtf/modules/newrelic/client"
)

type Client2 struct {
	applicationId int
	nrClient      *nr.Client
}

func NewClient(apiKey string, applicationId int) *Client2 {
	return &Client2{
		applicationId: applicationId,
		nrClient:      nr.NewClient(apiKey),
	}

}

func (client *Client2) Application() (*nr.Application, error) {

	application, err := client.nrClient.GetApplication(client.applicationId)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (client *Client2) Deployments() ([]nr.ApplicationDeployment, error) {

	opts := &nr.ApplicationDeploymentOptions{Page: 1}
	deployments, err := client.nrClient.GetApplicationDeployments(client.applicationId, opts)
	if err != nil {
		return nil, err
	}

	return deployments, nil
}
