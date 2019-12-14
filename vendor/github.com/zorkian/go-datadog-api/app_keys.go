/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

// APPKey represents an APP key
type APPKey struct {
	Owner *string `json:"owner,omitempty"`
	Name  *string `json:"name,omitemtpy"`
	Hash  *string `json:"hash,omitempty"`
}

// reqAPPKeys retrieves a slice of all APPKeys.
type reqAPPKeys struct {
	APPKeys []APPKey `json:"application_keys,omitempty"`
}

// reqAPPKey is similar to reqAPPKeys, but used for values returned by
// /v1/application_key/<somekey> which contain one object (not list) "application_key"
// (not "application_keys") containing the found key
type reqAPPKey struct {
	APPKey *APPKey `json:"application_key"`
}

// GetAPPKeys returns all APP keys or error on failure
func (client *Client) GetAPPKeys() ([]APPKey, error) {
	var out reqAPPKeys
	if err := client.doJsonRequest("GET", "/v1/application_key", nil, &out); err != nil {
		return nil, err
	}

	return out.APPKeys, nil
}

// GetAPPKey returns a single APP key or error on failure
func (client *Client) GetAPPKey(hash string) (*APPKey, error) {
	var out reqAPPKey
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/application_key/%s", hash), nil, &out); err != nil {
		return nil, err
	}

	return out.APPKey, nil
}

// CreateAPPKey creates an APP key from given name and fills the rest of its
// fields, or returns an error on failure
func (client *Client) CreateAPPKey(name string) (*APPKey, error) {
	toPost := struct {
		Name *string `json:"name,omitempty"`
	}{
		&name,
	}
	var out reqAPPKey
	if err := client.doJsonRequest("POST", "/v1/application_key", toPost, &out); err != nil {
		return nil, err
	}
	return out.APPKey, nil
}

// UpdateAPPKey updates given APP key (only Name can be updated), returns an error
func (client *Client) UpdateAPPKey(appkey *APPKey) error {
	out := reqAPPKey{APPKey: appkey}
	toPost := struct {
		Name *string `json:"name,omitempty"`
	}{
		appkey.Name,
	}
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/application_key/%s", *appkey.Hash), toPost, &out)
}

// DeleteAPPKey deletes APP key given by hash, returns an error
func (client *Client) DeleteAPPKey(hash string) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/application_key/%s", hash), nil, nil)
}
