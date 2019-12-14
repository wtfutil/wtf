// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package operations

import (
	"context"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
)

type Client interface {
	// Gets an operation from the the operationId using the given pluginId.
	GetOperation(context.Context, GetOperationArgs) (*Operation, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) Client {
	client := connection.GetClientByUrl(connection.BaseUrl)
	return &ClientImpl{
		Client: *client,
	}
}

// Gets an operation from the the operationId using the given pluginId.
func (client *ClientImpl) GetOperation(ctx context.Context, args GetOperationArgs) (*Operation, error) {
	routeValues := make(map[string]string)
	if args.OperationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.OperationId"}
	}
	routeValues["operationId"] = (*args.OperationId).String()

	queryParams := url.Values{}
	if args.PluginId != nil {
		queryParams.Add("pluginId", (*args.PluginId).String())
	}
	locationId, _ := uuid.Parse("9a1b74b4-2ca8-4a9f-8470-c2f2e6fdc949")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Operation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetOperation function
type GetOperationArgs struct {
	// (required) The ID for the operation.
	OperationId *uuid.UUID
	// (optional) The ID for the plugin.
	PluginId *uuid.UUID
}
