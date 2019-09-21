// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"context"
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

// Creates a new Azure DevOps connection instance using a personal access token.
func NewPatConnection(organizationUrl string, personalAccessToken string) *Connection {
	authorizationString := CreateBasicAuthHeaderValue("", personalAccessToken)
	organizationUrl = normalizeUrl(organizationUrl)
	return &Connection{
		AuthorizationString:     authorizationString,
		BaseUrl:                 organizationUrl,
		SuppressFedAuthRedirect: true,
	}
}

func NewAnonymousConnection(organizationUrl string) *Connection {
	organizationUrl = normalizeUrl(organizationUrl)
	return &Connection{
		BaseUrl:                 organizationUrl,
		SuppressFedAuthRedirect: true,
	}
}

type Connection struct {
	AuthorizationString     string
	BaseUrl                 string
	UserAgent               string
	SuppressFedAuthRedirect bool
	ForceMsaPassThrough     bool
	Timeout                 *time.Duration
	clientCache             map[string]Client
	clientCacheLock         sync.RWMutex
	resourceAreaCache       map[uuid.UUID]ResourceAreaInfo
	resourceAreaCacheLock   sync.RWMutex
}

func CreateBasicAuthHeaderValue(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func normalizeUrl(url string) string {
	return strings.ToLower(strings.TrimRight(url, "/"))
}

func (connection *Connection) GetClientByResourceAreaId(ctx context.Context, resourceAreaID uuid.UUID) (*Client, error) {
	resourceAreaInfo, err := connection.getResourceAreaInfo(ctx, resourceAreaID)
	if err != nil {
		return nil, err
	}
	var client *Client
	if resourceAreaInfo != nil {
		client = connection.GetClientByUrl(*resourceAreaInfo.LocationUrl)
	} else {
		// resourceAreaInfo will be nil for on prem servers
		client = connection.GetClientByUrl(connection.BaseUrl)
	}
	return client, nil
}

func (connection *Connection) GetClientByUrl(baseUrl string) *Client {
	normalizedUrl := normalizeUrl(baseUrl)
	azureDevOpsClient, ok := connection.getClientCacheEntry(normalizedUrl)
	if !ok {
		azureDevOpsClient = NewClient(connection, normalizedUrl)
		connection.setClientCacheEntry(normalizedUrl, azureDevOpsClient)
	}
	return azureDevOpsClient
}

func (connection *Connection) getResourceAreaInfo(ctx context.Context, resourceAreaId uuid.UUID) (*ResourceAreaInfo, error) {
	resourceAreaInfo, ok := connection.getResourceAreaCacheEntry(resourceAreaId)
	if !ok {
		client := connection.GetClientByUrl(connection.BaseUrl)
		resourceAreaInfos, err := client.GetResourceAreas(ctx)
		if err != nil {
			return nil, err
		}

		if len(*resourceAreaInfos) > 0 {
			for _, resourceEntry := range *resourceAreaInfos {
				connection.setResourceAreaCacheEntry(*resourceEntry.Id, &resourceEntry)
			}
			resourceAreaInfo, ok = connection.getResourceAreaCacheEntry(resourceAreaId)
		} else {
			// on prem servers return an empty list
			return nil, nil
		}
	}

	if ok {
		return resourceAreaInfo, nil
	}

	return nil, &ResourceAreaIdNotRegisteredError{resourceAreaId, connection.BaseUrl}
}

// Client Cache by Url
func (connection *Connection) getClientCacheEntry(url string) (*Client, bool) {
	if connection.clientCache == nil {
		return nil, false
	}
	connection.clientCacheLock.RLock()
	defer connection.clientCacheLock.RUnlock()
	client, ok := connection.clientCache[url]
	return &client, ok
}

func (connection *Connection) setClientCacheEntry(url string, client *Client) {
	connection.clientCacheLock.Lock()
	defer connection.clientCacheLock.Unlock()
	if connection.clientCache == nil {
		connection.clientCache = make(map[string]Client)
	}
	connection.clientCache[url] = *client
}

func (connection *Connection) getResourceAreaCacheEntry(resourceAreaId uuid.UUID) (*ResourceAreaInfo, bool) {
	if connection.resourceAreaCache == nil {
		return nil, false
	}
	connection.resourceAreaCacheLock.RLock()
	defer connection.resourceAreaCacheLock.RUnlock()
	resourceAreaInfo, ok := connection.resourceAreaCache[resourceAreaId]
	return &resourceAreaInfo, ok
}

func (connection *Connection) setResourceAreaCacheEntry(resourceAreaId uuid.UUID, resourceAreaInfo *ResourceAreaInfo) {
	connection.resourceAreaCacheLock.Lock()
	defer connection.resourceAreaCacheLock.Unlock()
	if connection.resourceAreaCache == nil {
		connection.resourceAreaCache = make(map[uuid.UUID]ResourceAreaInfo)
	}
	connection.resourceAreaCache[resourceAreaId] = *resourceAreaInfo
}

type ResourceAreaIdNotRegisteredError struct {
	ResourceAreaId uuid.UUID
	Url            string
}

func (e ResourceAreaIdNotRegisteredError) Error() string {
	return "API resource area Id " + e.ResourceAreaId.String() + " is not registered on " + e.Url + "."
}
