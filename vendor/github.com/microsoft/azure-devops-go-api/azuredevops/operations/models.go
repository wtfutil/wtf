// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package operations

import (
	"github.com/google/uuid"
)

// Contains information about the progress or result of an async operation.
type Operation struct {
	// Unique identifier for the operation.
	Id *uuid.UUID `json:"id,omitempty"`
	// Unique identifier for the plugin.
	PluginId *uuid.UUID `json:"pluginId,omitempty"`
	// The current status of the operation.
	Status *OperationStatus `json:"status,omitempty"`
	// URL to get the full operation object.
	Url *string `json:"url,omitempty"`
	// Links to other related objects.
	Links interface{} `json:"_links,omitempty"`
	// Detailed messaged about the status of an operation.
	DetailedMessage *string `json:"detailedMessage,omitempty"`
	// Result message for an operation.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// URL to the operation result.
	ResultUrl *OperationResultReference `json:"resultUrl,omitempty"`
}

// Reference for an async operation.
type OperationReference struct {
	// Unique identifier for the operation.
	Id *uuid.UUID `json:"id,omitempty"`
	// Unique identifier for the plugin.
	PluginId *uuid.UUID `json:"pluginId,omitempty"`
	// The current status of the operation.
	Status *OperationStatus `json:"status,omitempty"`
	// URL to get the full operation object.
	Url *string `json:"url,omitempty"`
}

type OperationResultReference struct {
	// URL to the operation result.
	ResultUrl *string `json:"resultUrl,omitempty"`
}

// The status of an operation.
type OperationStatus string

type operationStatusValuesType struct {
	NotSet     OperationStatus
	Queued     OperationStatus
	InProgress OperationStatus
	Cancelled  OperationStatus
	Succeeded  OperationStatus
	Failed     OperationStatus
}

var OperationStatusValues = operationStatusValuesType{
	// The operation does not have a status set.
	NotSet: "notSet",
	// The operation has been queued.
	Queued: "queued",
	// The operation is in progress.
	InProgress: "inProgress",
	// The operation was cancelled by the user.
	Cancelled: "cancelled",
	// The operation completed successfully.
	Succeeded: "succeeded",
	// The operation completed with a failure.
	Failed: "failed",
}
