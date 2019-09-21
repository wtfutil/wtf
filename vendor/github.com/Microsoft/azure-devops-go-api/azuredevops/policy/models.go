// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package policy

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

// The full policy configuration with settings.
type PolicyConfiguration struct {
	// The policy configuration ID.
	Id *int `json:"id,omitempty"`
	// The policy configuration type.
	Type *PolicyTypeRef `json:"type,omitempty"`
	// The URL where the policy configuration can be retrieved.
	Url *string `json:"url,omitempty"`
	// The policy configuration revision ID.
	Revision *int `json:"revision,omitempty"`
	// The links to other objects related to this object.
	Links interface{} `json:"_links,omitempty"`
	// A reference to the identity that created the policy.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// The date and time when the policy was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Indicates whether the policy is blocking.
	IsBlocking *bool `json:"isBlocking,omitempty"`
	// Indicates whether the policy has been (soft) deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Indicates whether the policy is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// The policy configuration settings.
	Settings interface{} `json:"settings,omitempty"`
}

// Policy configuration reference.
type PolicyConfigurationRef struct {
	// The policy configuration ID.
	Id *int `json:"id,omitempty"`
	// The policy configuration type.
	Type *PolicyTypeRef `json:"type,omitempty"`
	// The URL where the policy configuration can be retrieved.
	Url *string `json:"url,omitempty"`
}

// This record encapsulates the current state of a policy as it applies to one specific pull request. Each pull request has a unique PolicyEvaluationRecord for each pull request which the policy applies to.
type PolicyEvaluationRecord struct {
	// Links to other related objects
	Links interface{} `json:"_links,omitempty"`
	// A string which uniquely identifies the target of a policy evaluation.
	ArtifactId *string `json:"artifactId,omitempty"`
	// Time when this policy finished evaluating on this pull request.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Contains all configuration data for the policy which is being evaluated.
	Configuration *PolicyConfiguration `json:"configuration,omitempty"`
	// Internal context data of this policy evaluation.
	Context interface{} `json:"context,omitempty"`
	// Guid which uniquely identifies this evaluation record (one policy running on one pull request).
	EvaluationId *uuid.UUID `json:"evaluationId,omitempty"`
	// Time when this policy was first evaluated on this pull request.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// Status of the policy (Running, Approved, Failed, etc.)
	Status *PolicyEvaluationStatus `json:"status,omitempty"`
}

// Status of a policy which is running against a specific pull request.
type PolicyEvaluationStatus string

type policyEvaluationStatusValuesType struct {
	Queued        PolicyEvaluationStatus
	Running       PolicyEvaluationStatus
	Approved      PolicyEvaluationStatus
	Rejected      PolicyEvaluationStatus
	NotApplicable PolicyEvaluationStatus
	Broken        PolicyEvaluationStatus
}

var PolicyEvaluationStatusValues = policyEvaluationStatusValuesType{
	// The policy is either queued to run, or is waiting for some event before progressing.
	Queued: "queued",
	// The policy is currently running.
	Running: "running",
	// The policy has been fulfilled for this pull request.
	Approved: "approved",
	// The policy has rejected this pull request.
	Rejected: "rejected",
	// The policy does not apply to this pull request.
	NotApplicable: "notApplicable",
	// The policy has encountered an unexpected error.
	Broken: "broken",
}

// User-friendly policy type with description (used for querying policy types).
type PolicyType struct {
	// Display name of the policy type.
	DisplayName *string `json:"displayName,omitempty"`
	// The policy type ID.
	Id *uuid.UUID `json:"id,omitempty"`
	// The URL where the policy type can be retrieved.
	Url *string `json:"url,omitempty"`
	// The links to other objects related to this object.
	Links interface{} `json:"_links,omitempty"`
	// Detailed description of the policy type.
	Description *string `json:"description,omitempty"`
}

// Policy type reference.
type PolicyTypeRef struct {
	// Display name of the policy type.
	DisplayName *string `json:"displayName,omitempty"`
	// The policy type ID.
	Id *uuid.UUID `json:"id,omitempty"`
	// The URL where the policy type can be retrieved.
	Url *string `json:"url,omitempty"`
}

// A particular revision for a policy configuration.
type VersionedPolicyConfigurationRef struct {
	// The policy configuration ID.
	Id *int `json:"id,omitempty"`
	// The policy configuration type.
	Type *PolicyTypeRef `json:"type,omitempty"`
	// The URL where the policy configuration can be retrieved.
	Url *string `json:"url,omitempty"`
	// The policy configuration revision ID.
	Revision *int `json:"revision,omitempty"`
}
