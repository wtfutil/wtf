package newrelic

import (
	nr "github.com/yfronto/newrelic"
)

type Client struct {
	applicationId int
	nrClient      *nr.Client
}

func NewClient(apiKey string, applicationId int) *Client {
	return &Client{
		applicationId: applicationId,
		nrClient:      nr.NewClient(apiKey),
	}

}

func (client *Client) Application() (*nr.Application, error) {

	application, err := client.nrClient.GetApplication(client.applicationId)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (client *Client) Deployments() ([]nr.ApplicationDeployment, error) {

	opts := &nr.ApplicationDeploymentOptions{Page: 1}
	deployments, err := client.nrClient.GetApplicationDeployments(client.applicationId, opts)
	if err != nil {
		return nil, err
	}

	return deployments, nil
}
