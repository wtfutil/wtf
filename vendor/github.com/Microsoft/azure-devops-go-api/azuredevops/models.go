// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"encoding/json"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// ApiResourceLocation Information about the location of a REST API resource
type ApiResourceLocation struct {
	// Area name for this resource
	Area *string `json:"area,omitempty"`
	// Unique Identifier for this location
	Id *uuid.UUID `json:"id,omitempty"`
	// Maximum api version that this resource supports (current server version for this resource)
	MaxVersion *string `json:"maxVersion,omitempty"`
	// Minimum api version that this resource supports
	MinVersion *string `json:"minVersion,omitempty"`
	// The latest version of this resource location that is in "Release" (non-preview) mode
	ReleasedVersion *string `json:"releasedVersion,omitempty"`
	// Resource name
	ResourceName *string `json:"resourceName,omitempty"`
	// The current resource version supported by this resource location
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// This location's route template (templated relative path)
	RouteTemplate *string `json:"routeTemplate,omitempty"`
}

// WrappedImproperError
type WrappedImproperError struct {
	Count *int           `json:"count,omitEmpty"`
	Value *ImproperError `json:"value,omitEmpty"`
}

// ImproperError
type ImproperError struct {
	Message *string `json:"Message,omitEmpty"`
}

// KeyValuePair
type KeyValuePair struct {
	Key   *interface{} `json:"key,omitEmpty"`
	Value *interface{} `json:"value,omitEmpty"`
}

// ResourceAreaInfo
type ResourceAreaInfo struct {
	Id          *uuid.UUID `json:"id,omitempty"`
	LocationUrl *string    `json:"locationUrl,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type Time struct {
	Time time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	t2 := time.Time{}
	err := json.Unmarshal(b, &t2)

	// ignore errors for 0001-01-01T00:00:00 dates. The Azure DevOps service
	// returns default dates in a format that is invalid for a time.Time. The
	// correct value would have a 'z' at the end to represent utc. We are going
	// to ignore this error, and set the value to the default time.Time value.
	// https://github.com/microsoft/azure-devops-go-api/issues/17
	if err != nil {
		if parseError, ok := err.(*time.ParseError); ok && parseError.Value == "\"0001-01-01T00:00:00\"" {
			err = nil
		}
	}

	t.Time = t2
	return err
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func (t Time) String() string {
	return t.Time.String()
}

func (t Time) Equal(u Time) bool {
	return t.Time.Equal(u.Time)
}

// ServerSystemError
type ServerSystemError struct {
	ClassName      *string            `json:"className,omitempty"`
	InnerException *ServerSystemError `json:"innerException,omitempty"`
	Message        *string            `json:"message,omitempty"`
}

func (e ServerSystemError) Error() string {
	return *e.Message
}

// VssJsonCollectionWrapper -
type VssJsonCollectionWrapper struct {
	Count *int           `json:"count"`
	Value *[]interface{} `json:"value"`
}

// WrappedError
type WrappedError struct {
	ExceptionId      *string                 `json:"$id,omitempty"`
	InnerError       *WrappedError           `json:"innerException,omitempty"`
	Message          *string                 `json:"message,omitempty"`
	TypeName         *string                 `json:"typeName,omitempty"`
	TypeKey          *string                 `json:"typeKey,omitempty"`
	ErrorCode        *int                    `json:"errorCode,omitempty"`
	EventId          *int                    `json:"eventId,omitempty"`
	CustomProperties *map[string]interface{} `json:"customProperties,omitempty"`
	StatusCode       *int
}

func (e WrappedError) Error() string {
	if e.Message == nil {
		if e.StatusCode != nil {
			return "REST call returned status code " + strconv.Itoa(*e.StatusCode)
		}
		return ""
	}
	return *e.Message
}
