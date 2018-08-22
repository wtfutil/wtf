/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2018 by authors and contributors.
 */

package datadog

/*
	PagerDuty Integration
*/

type servicePD struct {
	ServiceName *string `json:"service_name"`
	ServiceKey  *string `json:"service_key"`
}

type integrationPD struct {
	Services  []servicePD `json:"services"`
	Subdomain *string     `json:"subdomain"`
	Schedules []string    `json:"schedules"`
	APIToken  *string     `json:"api_token"`
}

// ServicePDRequest defines the Services struct that is part of the IntegrationPDRequest.
type ServicePDRequest struct {
	ServiceName *string `json:"service_name"`
	ServiceKey  *string `json:"service_key"`
}

// IntegrationPDRequest defines the request payload for
// creating & updating Datadog-PagerDuty integration.
type IntegrationPDRequest struct {
	Services  []ServicePDRequest `json:"services,omitempty"`
	Subdomain *string            `json:"subdomain,omitempty"`
	Schedules []string           `json:"schedules,omitempty"`
	APIToken  *string            `json:"api_token,omitempty"`
	RunCheck  *bool              `json:"run_check,omitempty"`
}

// CreateIntegrationPD creates new PagerDuty Integrations.
// Use this if you want to setup the integration for the first time
// or to add more services/schedules.
func (client *Client) CreateIntegrationPD(pdIntegration *IntegrationPDRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/pagerduty", pdIntegration, nil)
}

// UpdateIntegrationPD updates the PagerDuty Integration.
// This will replace the existing values with the new values.
func (client *Client) UpdateIntegrationPD(pdIntegration *IntegrationPDRequest) error {
	return client.doJsonRequest("PUT", "/v1/integration/pagerduty", pdIntegration, nil)
}

// GetIntegrationPD gets all the PagerDuty Integrations from the system.
func (client *Client) GetIntegrationPD() (*integrationPD, error) {
	var out integrationPD
	if err := client.doJsonRequest("GET", "/v1/integration/pagerduty", nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// DeleteIntegrationPD removes the PagerDuty Integration from the system.
func (client *Client) DeleteIntegrationPD() error {
	return client.doJsonRequest("DELETE", "/v1/integration/pagerduty", nil, nil)
}

/*
	Slack Integration
*/

// ServiceHookSlackRequest defines the ServiceHooks struct that is part of the IntegrationSlackRequest.
type ServiceHookSlackRequest struct {
	Account *string `json:"account"`
	Url     *string `json:"url"`
}

// ChannelSlackRequest defines the Channels struct that is part of the IntegrationSlackRequest.
type ChannelSlackRequest struct {
	ChannelName             *string `json:"channel_name"`
	TransferAllUserComments *bool   `json:"transfer_all_user_comments,omitempty,string"`
	Account                 *string `json:"account"`
}

// IntegrationSlackRequest defines the request payload for
// creating & updating Datadog-Slack integration.
type IntegrationSlackRequest struct {
	ServiceHooks []ServiceHookSlackRequest `json:"service_hooks,omitempty"`
	Channels     []ChannelSlackRequest     `json:"channels,omitempty"`
	RunCheck     *bool                     `json:"run_check,omitempty,string"`
}

// CreateIntegrationSlack creates new Slack Integrations.
// Use this if you want to setup the integration for the first time
// or to add more channels.
func (client *Client) CreateIntegrationSlack(slackIntegration *IntegrationSlackRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/slack", slackIntegration, nil)
}

// UpdateIntegrationSlack updates the Slack Integration.
// This will replace the existing values with the new values.
func (client *Client) UpdateIntegrationSlack(slackIntegration *IntegrationSlackRequest) error {
	return client.doJsonRequest("PUT", "/v1/integration/slack", slackIntegration, nil)
}

// GetIntegrationSlack gets all the Slack Integrations from the system.
func (client *Client) GetIntegrationSlack() (*IntegrationSlackRequest, error) {
	var out IntegrationSlackRequest
	if err := client.doJsonRequest("GET", "/v1/integration/slack", nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// DeleteIntegrationSlack removes the Slack Integration from the system.
func (client *Client) DeleteIntegrationSlack() error {
	return client.doJsonRequest("DELETE", "/v1/integration/slack", nil, nil)
}

/*
	AWS Integration
*/

// IntegrationAWSAccount defines the request payload for
// creating & updating Datadog-AWS integration.
type IntegrationAWSAccount struct {
	AccountID                     *string         `json:"account_id"`
	RoleName                      *string         `json:"role_name"`
	FilterTags                    []string        `json:"filter_tags"`
	HostTags                      []string        `json:"host_tags"`
	AccountSpecificNamespaceRules map[string]bool `json:"account_specific_namespace_rules"`
}

// IntegrationAWSAccountCreateResponse defines the response payload for
// creating & updating Datadog-AWS integration.
type IntegrationAWSAccountCreateResponse struct {
	ExternalID string `json:"external_id"`
}

type IntegrationAWSAccountGetResponse struct {
	Accounts []IntegrationAWSAccount `json:"accounts"`
}

type IntegrationAWSAccountDeleteRequest struct {
	AccountID *string `json:"account_id"`
	RoleName  *string `json:"role_name"`
}

// CreateIntegrationAWS adds a new AWS Account in the AWS Integrations.
// Use this if you want to setup the integration for the first time
// or to add more accounts.
func (client *Client) CreateIntegrationAWS(awsAccount *IntegrationAWSAccount) (*IntegrationAWSAccountCreateResponse, error) {
	var out IntegrationAWSAccountCreateResponse
	if err := client.doJsonRequest("POST", "/v1/integration/aws", awsAccount, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// GetIntegrationAWS gets all the AWS Accounts in the AWS Integrations from Datadog.
func (client *Client) GetIntegrationAWS() (*[]IntegrationAWSAccount, error) {
	var response IntegrationAWSAccountGetResponse
	if err := client.doJsonRequest("GET", "/v1/integration/aws", nil, &response); err != nil {
		return nil, err
	}

	return &response.Accounts, nil
}

// DeleteIntegrationAWS removes a specific AWS Account from the AWS Integration.
func (client *Client) DeleteIntegrationAWS(awsAccount *IntegrationAWSAccountDeleteRequest) error {
	return client.doJsonRequest("DELETE", "/v1/integration/aws", awsAccount, nil)
}
