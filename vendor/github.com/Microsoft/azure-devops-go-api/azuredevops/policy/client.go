// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package policy

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("fb13a388-40dd-4a04-b530-013a739c72ef")

type Client struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (*Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: *client,
	}, nil
}

// Create a policy configuration of a given policy type.
func (client *Client) CreatePolicyConfiguration(ctx context.Context, args CreatePolicyConfigurationArgs) (*PolicyConfiguration, error) {
	if args.Configuration == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Configuration"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ConfigurationId != nil {
		routeValues["configurationId"] = strconv.Itoa(*args.ConfigurationId)
	}

	body, marshalErr := json.Marshal(*args.Configuration)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePolicyConfiguration function
type CreatePolicyConfigurationArgs struct {
	// (required) The policy configuration to create.
	Configuration *PolicyConfiguration
	// (required) Project ID or project name
	Project *string
	// (optional)
	ConfigurationId *int
}

// Delete a policy configuration by its ID.
func (client *Client) DeletePolicyConfiguration(ctx context.Context, args DeletePolicyConfigurationArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ConfigurationId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ConfigurationId"}
	}
	routeValues["configurationId"] = strconv.Itoa(*args.ConfigurationId)

	locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePolicyConfiguration function
type DeletePolicyConfigurationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the policy configuration to delete.
	ConfigurationId *int
}

// Get a policy configuration by its ID.
func (client *Client) GetPolicyConfiguration(ctx context.Context, args GetPolicyConfigurationArgs) (*PolicyConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ConfigurationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ConfigurationId"}
	}
	routeValues["configurationId"] = strconv.Itoa(*args.ConfigurationId)

	locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyConfiguration function
type GetPolicyConfigurationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the policy configuration
	ConfigurationId *int
}

// Get a list of policy configurations in a project.
func (client *Client) GetPolicyConfigurations(ctx context.Context, args GetPolicyConfigurationsArgs) (*GetPolicyConfigurationsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Scope != nil {
		queryParams.Add("scope", *args.Scope)
	}
	if args.PolicyType != nil {
		queryParams.Add("policyType", (*args.PolicyType).String())
	}
	locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetPolicyConfigurationsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetPolicyConfigurations function
type GetPolicyConfigurationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) [Provided for legacy reasons] The scope on which a subset of policies is defined.
	Scope *string
	// (optional) Filter returned policies to only this type
	PolicyType *uuid.UUID
}

// Return type for the GetPolicyConfigurations function
type GetPolicyConfigurationsResponseValue struct {
	Value []PolicyConfiguration
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// Update a policy configuration by its ID.
func (client *Client) UpdatePolicyConfiguration(ctx context.Context, args UpdatePolicyConfigurationArgs) (*PolicyConfiguration, error) {
	if args.Configuration == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Configuration"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ConfigurationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ConfigurationId"}
	}
	routeValues["configurationId"] = strconv.Itoa(*args.ConfigurationId)

	body, marshalErr := json.Marshal(*args.Configuration)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdatePolicyConfiguration function
type UpdatePolicyConfigurationArgs struct {
	// (required) The policy configuration to update.
	Configuration *PolicyConfiguration
	// (required) Project ID or project name
	Project *string
	// (required) ID of the existing policy configuration to be updated.
	ConfigurationId *int
}

// [Preview API] Gets the present evaluation state of a policy.
func (client *Client) GetPolicyEvaluation(ctx context.Context, args GetPolicyEvaluationArgs) (*PolicyEvaluationRecord, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.EvaluationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.EvaluationId"}
	}
	routeValues["evaluationId"] = (*args.EvaluationId).String()

	locationId, _ := uuid.Parse("46aecb7a-5d2c-4647-897b-0209505a9fe4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyEvaluationRecord
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyEvaluation function
type GetPolicyEvaluationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the policy evaluation to be retrieved.
	EvaluationId *uuid.UUID
}

// [Preview API] Requeue the policy evaluation.
func (client *Client) RequeuePolicyEvaluation(ctx context.Context, args RequeuePolicyEvaluationArgs) (*PolicyEvaluationRecord, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.EvaluationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.EvaluationId"}
	}
	routeValues["evaluationId"] = (*args.EvaluationId).String()

	locationId, _ := uuid.Parse("46aecb7a-5d2c-4647-897b-0209505a9fe4")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyEvaluationRecord
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the RequeuePolicyEvaluation function
type RequeuePolicyEvaluationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the policy evaluation to be retrieved.
	EvaluationId *uuid.UUID
}

// [Preview API] Retrieves a list of all the policy evaluation statuses for a specific pull request.
func (client *Client) GetPolicyEvaluations(ctx context.Context, args GetPolicyEvaluationsArgs) (*[]PolicyEvaluationRecord, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.ArtifactId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "artifactId"}
	}
	queryParams.Add("artifactId", *args.ArtifactId)
	if args.IncludeNotApplicable != nil {
		queryParams.Add("includeNotApplicable", strconv.FormatBool(*args.IncludeNotApplicable))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("c23ddff5-229c-4d04-a80b-0fdce9f360c8")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []PolicyEvaluationRecord
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyEvaluations function
type GetPolicyEvaluationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) A string which uniquely identifies the target of a policy evaluation.
	ArtifactId *string
	// (optional) Some policies might determine that they do not apply to a specific pull request. Setting this parameter to true will return evaluation records even for policies which don't apply to this pull request.
	IncludeNotApplicable *bool
	// (optional) The number of policy evaluation records to retrieve.
	Top *int
	// (optional) The number of policy evaluation records to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
	Skip *int
}

// Retrieve a specific revision of a given policy by ID.
func (client *Client) GetPolicyConfigurationRevision(ctx context.Context, args GetPolicyConfigurationRevisionArgs) (*PolicyConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ConfigurationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ConfigurationId"}
	}
	routeValues["configurationId"] = strconv.Itoa(*args.ConfigurationId)
	if args.RevisionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RevisionId"}
	}
	routeValues["revisionId"] = strconv.Itoa(*args.RevisionId)

	locationId, _ := uuid.Parse("fe1e68a2-60d3-43cb-855b-85e41ae97c95")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyConfigurationRevision function
type GetPolicyConfigurationRevisionArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The policy configuration ID.
	ConfigurationId *int
	// (required) The revision ID.
	RevisionId *int
}

// Retrieve all revisions for a given policy.
func (client *Client) GetPolicyConfigurationRevisions(ctx context.Context, args GetPolicyConfigurationRevisionsArgs) (*[]PolicyConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ConfigurationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ConfigurationId"}
	}
	routeValues["configurationId"] = strconv.Itoa(*args.ConfigurationId)

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("fe1e68a2-60d3-43cb-855b-85e41ae97c95")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []PolicyConfiguration
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyConfigurationRevisions function
type GetPolicyConfigurationRevisionsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The policy configuration ID.
	ConfigurationId *int
	// (optional) The number of revisions to retrieve.
	Top *int
	// (optional) The number of revisions to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
	Skip *int
}

// Retrieve a specific policy type by ID.
func (client *Client) GetPolicyType(ctx context.Context, args GetPolicyTypeArgs) (*PolicyType, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TypeId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TypeId"}
	}
	routeValues["typeId"] = (*args.TypeId).String()

	locationId, _ := uuid.Parse("44096322-2d3d-466a-bb30-d1b7de69f61f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PolicyType
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyType function
type GetPolicyTypeArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The policy ID.
	TypeId *uuid.UUID
}

// Retrieve all available policy types.
func (client *Client) GetPolicyTypes(ctx context.Context, args GetPolicyTypesArgs) (*[]PolicyType, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("44096322-2d3d-466a-bb30-d1b7de69f61f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []PolicyType
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPolicyTypes function
type GetPolicyTypesArgs struct {
	// (required) Project ID or project name
	Project *string
}
