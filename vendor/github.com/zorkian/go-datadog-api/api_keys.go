/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"fmt"
	"time"
)

var createdTimeLayout = "2006-01-02 15:04:05"

// APIKey represents and API key
type APIKey struct {
	CreatedBy *string    `json:"created_by,omitempty"`
	Name      *string    `json:"name,omitemtpy"`
	Key       *string    `json:"key,omitempty"`
	Created   *time.Time `json:"created,omitempty"`
}

// reqAPIKeys retrieves a slice of all APIKeys.
type reqAPIKeys struct {
	APIKeys []APIKey `json:"api_keys,omitempty"`
}

// reqAPIKey is similar to reqAPIKeys, but used for values returned by /v1/api_key/<somekey>
// which contain one object (not list) "api_key" (not "api_keys") containing the found key
type reqAPIKey struct {
	APIKey *APIKey `json:"api_key"`
}

// MarshalJSON is a custom method for handling datetime marshalling
func (k APIKey) MarshalJSON() ([]byte, error) {
	// Approach for custom (un)marshalling borrowed from http://choly.ca/post/go-json-marshalling/
	type Alias APIKey
	return json.Marshal(&struct {
		Created *string `json:"created,omitempty"`
		*Alias
	}{
		Created: String(k.Created.Format(createdTimeLayout)),
		Alias:   (*Alias)(&k),
	})
}

// UnmarshalJSON is a custom method for handling datetime unmarshalling
func (k *APIKey) UnmarshalJSON(data []byte) error {
	type Alias APIKey
	aux := &struct {
		Created *string `json:"created,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(k),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if created, err := time.Parse(createdTimeLayout, *aux.Created); err != nil {
		return err
	} else {
		k.Created = &created
	}

	return nil
}

// GetAPIKeys returns all API keys or error on failure
func (client *Client) GetAPIKeys() ([]APIKey, error) {
	var out reqAPIKeys
	if err := client.doJsonRequest("GET", "/v1/api_key", nil, &out); err != nil {
		return nil, err
	}

	return out.APIKeys, nil
}

// GetAPIKey returns a single API key or error on failure
func (client *Client) GetAPIKey(key string) (*APIKey, error) {
	var out reqAPIKey
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/api_key/%s", key), nil, &out); err != nil {
		return nil, err
	}

	return out.APIKey, nil
}

// CreateAPIKey creates an API key from given struct and fills the rest of its
// fields, or returns an error on failure
func (client *Client) CreateAPIKey(name string) (*APIKey, error) {
	toPost := struct {
		Name *string `json:"name,omitempty"`
	}{
		&name,
	}
	var out reqAPIKey
	if err := client.doJsonRequest("POST", "/v1/api_key", toPost, &out); err != nil {
		return nil, err
	}
	return out.APIKey, nil
}

// UpdateAPIKey updates given API key (only Name can be updated), returns an error
func (client *Client) UpdateAPIKey(apikey *APIKey) error {
	out := reqAPIKey{APIKey: apikey}
	toPost := struct {
		Name *string `json:"name,omitempty"`
	}{
		apikey.Name,
	}
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/api_key/%s", *apikey.Key), toPost, &out)
}

// DeleteAPIKey deletes API key given by key, returns an error
func (client *Client) DeleteAPIKey(key string) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/api_key/%s", key), nil, nil)
}
