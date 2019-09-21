// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package build

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/distributedtaskcommon"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/microsoft/azure-devops-go-api/azuredevops/test"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

// Represents a queue for running builds.
type AgentPoolQueue struct {
	Links interface{} `json:"_links,omitempty"`
	// The ID of the queue.
	Id *int `json:"id,omitempty"`
	// The name of the queue.
	Name *string `json:"name,omitempty"`
	// The pool used by this queue.
	Pool *TaskAgentPoolReference `json:"pool,omitempty"`
	// The full http link to the resource.
	Url *string `json:"url,omitempty"`
}

// Represents a reference to an agent queue.
type AgentPoolQueueReference struct {
	// An alias to be used when referencing the resource.
	Alias *string `json:"alias,omitempty"`
	// The ID of the queue.
	Id *int `json:"id,omitempty"`
}

// Describes how a phase should run against an agent queue.
type AgentPoolQueueTarget struct {
	// The type of the target.
	Type *int `json:"type,omitempty"`
	// Agent specification of the target.
	AgentSpecification *AgentSpecification `json:"agentSpecification,omitempty"`
	// Enables scripts and other processes launched while executing phase to access the OAuth token
	AllowScriptsAuthAccessOption *bool          `json:"allowScriptsAuthAccessOption,omitempty"`
	Demands                      *[]interface{} `json:"demands,omitempty"`
	// The execution options.
	ExecutionOptions *AgentTargetExecutionOptions `json:"executionOptions,omitempty"`
	// The queue.
	Queue *AgentPoolQueue `json:"queue,omitempty"`
}

// Specification of the agent defined by the pool provider.
type AgentSpecification struct {
	// Agent specification unique identifier.
	Identifier *string `json:"identifier,omitempty"`
}

type AgentStatus string

type agentStatusValuesType struct {
	Unavailable AgentStatus
	Available   AgentStatus
	Offline     AgentStatus
}

var AgentStatusValues = agentStatusValuesType{
	// Indicates that the build agent cannot be contacted.
	Unavailable: "unavailable",
	// Indicates that the build agent is currently available.
	Available: "available",
	// Indicates that the build agent has taken itself offline.
	Offline: "offline",
}

// Additional options for running phases against an agent queue.
type AgentTargetExecutionOptions struct {
	// Indicates the type of execution options.
	Type *int `json:"type,omitempty"`
}

type ArtifactResource struct {
	Links interface{} `json:"_links,omitempty"`
	// Type-specific data about the artifact.
	Data *string `json:"data,omitempty"`
	// A link to download the resource.
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// Type-specific properties of the artifact.
	Properties *map[string]string `json:"properties,omitempty"`
	// The type of the resource: File container, version control folder, UNC path, etc.
	Type *string `json:"type,omitempty"`
	// The full http link to the resource.
	Url *string `json:"url,omitempty"`
}

// Represents an attachment to a build.
type Attachment struct {
	Links interface{} `json:"_links,omitempty"`
	// The name of the attachment.
	Name *string `json:"name,omitempty"`
}

type AuditAction string

type auditActionValuesType struct {
	Add    AuditAction
	Update AuditAction
	Delete AuditAction
}

var AuditActionValues = auditActionValuesType{
	Add:    "add",
	Update: "update",
	Delete: "delete",
}

// Data representation of a build.
type Build struct {
	Links interface{} `json:"_links,omitempty"`
	// The agent specification for the build.
	AgentSpecification *AgentSpecification `json:"agentSpecification,omitempty"`
	// The build number/name of the build.
	BuildNumber *string `json:"buildNumber,omitempty"`
	// The build number revision.
	BuildNumberRevision *int `json:"buildNumberRevision,omitempty"`
	// The build controller. This is only set if the definition type is Xaml.
	Controller *BuildController `json:"controller,omitempty"`
	// The definition associated with the build.
	Definition *DefinitionReference `json:"definition,omitempty"`
	// Indicates whether the build has been deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// The identity of the process or person that deleted the build.
	DeletedBy *webapi.IdentityRef `json:"deletedBy,omitempty"`
	// The date the build was deleted.
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
	// The description of how the build was deleted.
	DeletedReason *string `json:"deletedReason,omitempty"`
	// A list of demands that represents the agent capabilities required by this build.
	Demands *[]interface{} `json:"demands,omitempty"`
	// The time that the build was completed.
	FinishTime *azuredevops.Time `json:"finishTime,omitempty"`
	// The ID of the build.
	Id *int `json:"id,omitempty"`
	// Indicates whether the build should be skipped by retention policies.
	KeepForever *bool `json:"keepForever,omitempty"`
	// The identity representing the process or person that last changed the build.
	LastChangedBy *webapi.IdentityRef `json:"lastChangedBy,omitempty"`
	// The date the build was last changed.
	LastChangedDate *azuredevops.Time `json:"lastChangedDate,omitempty"`
	// Information about the build logs.
	Logs *BuildLogReference `json:"logs,omitempty"`
	// The orchestration plan for the build.
	OrchestrationPlan *TaskOrchestrationPlanReference `json:"orchestrationPlan,omitempty"`
	// The parameters for the build.
	Parameters *string `json:"parameters,omitempty"`
	// Orchestration plans associated with the build (build, cleanup)
	Plans *[]TaskOrchestrationPlanReference `json:"plans,omitempty"`
	// The build's priority.
	Priority *QueuePriority `json:"priority,omitempty"`
	// The team project.
	Project    *core.TeamProjectReference `json:"project,omitempty"`
	Properties interface{}                `json:"properties,omitempty"`
	// The quality of the xaml build (good, bad, etc.)
	Quality *string `json:"quality,omitempty"`
	// The queue. This is only set if the definition type is Build.
	Queue *AgentPoolQueue `json:"queue,omitempty"`
	// Additional options for queueing the build.
	QueueOptions *QueueOptions `json:"queueOptions,omitempty"`
	// The current position of the build in the queue.
	QueuePosition *int `json:"queuePosition,omitempty"`
	// The time that the build was queued.
	QueueTime *azuredevops.Time `json:"queueTime,omitempty"`
	// The reason that the build was created.
	Reason *BuildReason `json:"reason,omitempty"`
	// The repository.
	Repository *BuildRepository `json:"repository,omitempty"`
	// The identity that queued the build.
	RequestedBy *webapi.IdentityRef `json:"requestedBy,omitempty"`
	// The identity on whose behalf the build was queued.
	RequestedFor *webapi.IdentityRef `json:"requestedFor,omitempty"`
	// The build result.
	Result *BuildResult `json:"result,omitempty"`
	// Indicates whether the build is retained by a release.
	RetainedByRelease *bool `json:"retainedByRelease,omitempty"`
	// The source branch.
	SourceBranch *string `json:"sourceBranch,omitempty"`
	// The source version.
	SourceVersion *string `json:"sourceVersion,omitempty"`
	// The time that the build was started.
	StartTime *azuredevops.Time `json:"startTime,omitempty"`
	// The status of the build.
	Status *BuildStatus `json:"status,omitempty"`
	Tags   *[]string    `json:"tags,omitempty"`
	// The build that triggered this build via a Build completion trigger.
	TriggeredByBuild *Build `json:"triggeredByBuild,omitempty"`
	// Sourceprovider-specific information about what triggered the build
	TriggerInfo *map[string]string `json:"triggerInfo,omitempty"`
	// The URI of the build.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the build.
	Url               *string                         `json:"url,omitempty"`
	ValidationResults *[]BuildRequestValidationResult `json:"validationResults,omitempty"`
}

type BuildAgent struct {
	BuildDirectory   *string                       `json:"buildDirectory,omitempty"`
	Controller       *XamlBuildControllerReference `json:"controller,omitempty"`
	CreatedDate      *azuredevops.Time             `json:"createdDate,omitempty"`
	Description      *string                       `json:"description,omitempty"`
	Enabled          *bool                         `json:"enabled,omitempty"`
	Id               *int                          `json:"id,omitempty"`
	MessageQueueUrl  *string                       `json:"messageQueueUrl,omitempty"`
	Name             *string                       `json:"name,omitempty"`
	ReservedForBuild *string                       `json:"reservedForBuild,omitempty"`
	Server           *XamlBuildServerReference     `json:"server,omitempty"`
	Status           *AgentStatus                  `json:"status,omitempty"`
	StatusMessage    *string                       `json:"statusMessage,omitempty"`
	UpdatedDate      *azuredevops.Time             `json:"updatedDate,omitempty"`
	Uri              *string                       `json:"uri,omitempty"`
	Url              *string                       `json:"url,omitempty"`
}

type BuildAgentReference struct {
	// Id of the resource
	Id *int `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

// Represents an artifact produced by a build.
type BuildArtifact struct {
	// The artifact ID.
	Id *int `json:"id,omitempty"`
	// The name of the artifact.
	Name *string `json:"name,omitempty"`
	// The actual resource.
	Resource *ArtifactResource `json:"resource,omitempty"`
	// The artifact source, which will be the ID of the job that produced this artifact.
	Source *string `json:"source,omitempty"`
}

// Represents the desired scope of authorization for a build.
type BuildAuthorizationScope string

type buildAuthorizationScopeValuesType struct {
	ProjectCollection BuildAuthorizationScope
	Project           BuildAuthorizationScope
}

var BuildAuthorizationScopeValues = buildAuthorizationScopeValuesType{
	// The identity used should have build service account permissions scoped to the project collection. This is useful when resources for a single build are spread across multiple projects.
	ProjectCollection: "projectCollection",
	// The identity used should have build service account permissions scoped to the project in which the build definition resides. This is useful for isolation of build jobs to a particular team project to avoid any unintentional escalation of privilege attacks during a build.
	Project: "project",
}

// Represents a build badge.
type BuildBadge struct {
	// The ID of the build represented by this badge.
	BuildId *int `json:"buildId,omitempty"`
	// A link to the SVG resource.
	ImageUrl *string `json:"imageUrl,omitempty"`
}

type BuildCompletedEvent struct {
	BuildId *int   `json:"buildId,omitempty"`
	Build   *Build `json:"build,omitempty"`
	// Changes associated with a build used for build notifications
	Changes *[]Change `json:"changes,omitempty"`
	// Pull request for the build used for build notifications
	PullRequest *PullRequest `json:"pullRequest,omitempty"`
	// Test results associated with a build used for build notifications
	TestResults *test.AggregatedResultsAnalysis `json:"testResults,omitempty"`
	// Timeline records associated with a build used for build notifications
	TimelineRecords *[]TimelineRecord `json:"timelineRecords,omitempty"`
	// Work items associated with a build used for build notifications
	WorkItems *[]git.AssociatedWorkItem `json:"workItems,omitempty"`
}

// Represents a build completion trigger.
type BuildCompletionTrigger struct {
	BranchFilters *[]string `json:"branchFilters,omitempty"`
	// A reference to the definition that should trigger builds for this definition.
	Definition              *DefinitionReference `json:"definition,omitempty"`
	RequiresSuccessfulBuild *bool                `json:"requiresSuccessfulBuild,omitempty"`
}

type BuildController struct {
	// Id of the resource
	Id *int `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url   *string     `json:"url,omitempty"`
	Links interface{} `json:"_links,omitempty"`
	// The date the controller was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The description of the controller.
	Description *string `json:"description,omitempty"`
	// Indicates whether the controller is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The status of the controller.
	Status *ControllerStatus `json:"status,omitempty"`
	// The date the controller was last updated.
	UpdatedDate *azuredevops.Time `json:"updatedDate,omitempty"`
	// The controller's URI.
	Uri *string `json:"uri,omitempty"`
}

// Represents a build definition.
type BuildDefinition struct {
	// The date this version of the definition was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The ID of the referenced definition.
	Id *int `json:"id,omitempty"`
	// The name of the referenced definition.
	Name *string `json:"name,omitempty"`
	// The folder path of the definition.
	Path *string `json:"path,omitempty"`
	// A reference to the project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// A value that indicates whether builds can be queued against this definition.
	QueueStatus *DefinitionQueueStatus `json:"queueStatus,omitempty"`
	// The definition revision number.
	Revision *int `json:"revision,omitempty"`
	// The type of the definition.
	Type *DefinitionType `json:"type,omitempty"`
	// The definition's URI.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the definition.
	Url   *string     `json:"url,omitempty"`
	Links interface{} `json:"_links,omitempty"`
	// The author of the definition.
	AuthoredBy *webapi.IdentityRef `json:"authoredBy,omitempty"`
	// A reference to the definition that this definition is a draft of, if this is a draft definition.
	DraftOf *DefinitionReference `json:"draftOf,omitempty"`
	// The list of drafts associated with this definition, if this is not a draft definition.
	Drafts               *[]DefinitionReference `json:"drafts,omitempty"`
	LatestBuild          *Build                 `json:"latestBuild,omitempty"`
	LatestCompletedBuild *Build                 `json:"latestCompletedBuild,omitempty"`
	Metrics              *[]BuildMetric         `json:"metrics,omitempty"`
	// The quality of the definition document (draft, etc.)
	Quality *DefinitionQuality `json:"quality,omitempty"`
	// The default queue for builds run against this definition.
	Queue *AgentPoolQueue `json:"queue,omitempty"`
	// Indicates whether badges are enabled for this definition.
	BadgeEnabled *bool `json:"badgeEnabled,omitempty"`
	// The build number format.
	BuildNumberFormat *string `json:"buildNumberFormat,omitempty"`
	// A save-time comment for the definition.
	Comment *string        `json:"comment,omitempty"`
	Demands *[]interface{} `json:"demands,omitempty"`
	// The description.
	Description *string `json:"description,omitempty"`
	// The drop location for the definition.
	DropLocation *string `json:"dropLocation,omitempty"`
	// The job authorization scope for builds queued against this definition.
	JobAuthorizationScope *BuildAuthorizationScope `json:"jobAuthorizationScope,omitempty"`
	// The job cancel timeout (in minutes) for builds cancelled by user for this definition.
	JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
	// The job execution timeout (in minutes) for builds queued against this definition.
	JobTimeoutInMinutes *int           `json:"jobTimeoutInMinutes,omitempty"`
	Options             *[]BuildOption `json:"options,omitempty"`
	// The build process.
	Process interface{} `json:"process,omitempty"`
	// The process parameters for this definition.
	ProcessParameters *distributedtaskcommon.ProcessParameters `json:"processParameters,omitempty"`
	Properties        interface{}                              `json:"properties,omitempty"`
	// The repository.
	Repository     *BuildRepository                    `json:"repository,omitempty"`
	RetentionRules *[]RetentionPolicy                  `json:"retentionRules,omitempty"`
	Tags           *[]string                           `json:"tags,omitempty"`
	Triggers       *[]interface{}                      `json:"triggers,omitempty"`
	VariableGroups *[]VariableGroup                    `json:"variableGroups,omitempty"`
	Variables      *map[string]BuildDefinitionVariable `json:"variables,omitempty"`
}

// For back-compat with extensions that use the old Steps format instead of Process and Phases
type BuildDefinition3_2 struct {
	// The date this version of the definition was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The ID of the referenced definition.
	Id *int `json:"id,omitempty"`
	// The name of the referenced definition.
	Name *string `json:"name,omitempty"`
	// The folder path of the definition.
	Path *string `json:"path,omitempty"`
	// A reference to the project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// A value that indicates whether builds can be queued against this definition.
	QueueStatus *DefinitionQueueStatus `json:"queueStatus,omitempty"`
	// The definition revision number.
	Revision *int `json:"revision,omitempty"`
	// The type of the definition.
	Type *DefinitionType `json:"type,omitempty"`
	// The definition's URI.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the definition.
	Url   *string     `json:"url,omitempty"`
	Links interface{} `json:"_links,omitempty"`
	// The author of the definition.
	AuthoredBy *webapi.IdentityRef `json:"authoredBy,omitempty"`
	// A reference to the definition that this definition is a draft of, if this is a draft definition.
	DraftOf *DefinitionReference `json:"draftOf,omitempty"`
	// The list of drafts associated with this definition, if this is not a draft definition.
	Drafts  *[]DefinitionReference `json:"drafts,omitempty"`
	Metrics *[]BuildMetric         `json:"metrics,omitempty"`
	// The quality of the definition document (draft, etc.)
	Quality *DefinitionQuality `json:"quality,omitempty"`
	// The default queue for builds run against this definition.
	Queue *AgentPoolQueue `json:"queue,omitempty"`
	// Indicates whether badges are enabled for this definition
	BadgeEnabled *bool                  `json:"badgeEnabled,omitempty"`
	Build        *[]BuildDefinitionStep `json:"build,omitempty"`
	// The build number format
	BuildNumberFormat *string `json:"buildNumberFormat,omitempty"`
	// The comment entered when saving the definition
	Comment *string        `json:"comment,omitempty"`
	Demands *[]interface{} `json:"demands,omitempty"`
	// The description
	Description *string `json:"description,omitempty"`
	// The drop location for the definition
	DropLocation *string `json:"dropLocation,omitempty"`
	// The job authorization scope for builds which are queued against this definition
	JobAuthorizationScope *BuildAuthorizationScope `json:"jobAuthorizationScope,omitempty"`
	// The job cancel timeout in minutes for builds which are cancelled by user for this definition
	JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
	// The job execution timeout in minutes for builds which are queued against this definition
	JobTimeoutInMinutes  *int           `json:"jobTimeoutInMinutes,omitempty"`
	LatestBuild          *Build         `json:"latestBuild,omitempty"`
	LatestCompletedBuild *Build         `json:"latestCompletedBuild,omitempty"`
	Options              *[]BuildOption `json:"options,omitempty"`
	// Process Parameters
	ProcessParameters *distributedtaskcommon.ProcessParameters `json:"processParameters,omitempty"`
	Properties        interface{}                              `json:"properties,omitempty"`
	// The repository
	Repository     *BuildRepository                    `json:"repository,omitempty"`
	RetentionRules *[]RetentionPolicy                  `json:"retentionRules,omitempty"`
	Tags           *[]string                           `json:"tags,omitempty"`
	Triggers       *[]interface{}                      `json:"triggers,omitempty"`
	Variables      *map[string]BuildDefinitionVariable `json:"variables,omitempty"`
}

// Represents a reference to a build definition.
type BuildDefinitionReference struct {
	// The date this version of the definition was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The ID of the referenced definition.
	Id *int `json:"id,omitempty"`
	// The name of the referenced definition.
	Name *string `json:"name,omitempty"`
	// The folder path of the definition.
	Path *string `json:"path,omitempty"`
	// A reference to the project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// A value that indicates whether builds can be queued against this definition.
	QueueStatus *DefinitionQueueStatus `json:"queueStatus,omitempty"`
	// The definition revision number.
	Revision *int `json:"revision,omitempty"`
	// The type of the definition.
	Type *DefinitionType `json:"type,omitempty"`
	// The definition's URI.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the definition.
	Url   *string     `json:"url,omitempty"`
	Links interface{} `json:"_links,omitempty"`
	// The author of the definition.
	AuthoredBy *webapi.IdentityRef `json:"authoredBy,omitempty"`
	// A reference to the definition that this definition is a draft of, if this is a draft definition.
	DraftOf *DefinitionReference `json:"draftOf,omitempty"`
	// The list of drafts associated with this definition, if this is not a draft definition.
	Drafts               *[]DefinitionReference `json:"drafts,omitempty"`
	LatestBuild          *Build                 `json:"latestBuild,omitempty"`
	LatestCompletedBuild *Build                 `json:"latestCompletedBuild,omitempty"`
	Metrics              *[]BuildMetric         `json:"metrics,omitempty"`
	// The quality of the definition document (draft, etc.)
	Quality *DefinitionQuality `json:"quality,omitempty"`
	// The default queue for builds run against this definition.
	Queue *AgentPoolQueue `json:"queue,omitempty"`
}

// For back-compat with extensions that use the old Steps format instead of Process and Phases
type BuildDefinitionReference3_2 struct {
	// The date this version of the definition was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The ID of the referenced definition.
	Id *int `json:"id,omitempty"`
	// The name of the referenced definition.
	Name *string `json:"name,omitempty"`
	// The folder path of the definition.
	Path *string `json:"path,omitempty"`
	// A reference to the project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// A value that indicates whether builds can be queued against this definition.
	QueueStatus *DefinitionQueueStatus `json:"queueStatus,omitempty"`
	// The definition revision number.
	Revision *int `json:"revision,omitempty"`
	// The type of the definition.
	Type *DefinitionType `json:"type,omitempty"`
	// The definition's URI.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the definition.
	Url   *string     `json:"url,omitempty"`
	Links interface{} `json:"_links,omitempty"`
	// The author of the definition.
	AuthoredBy *webapi.IdentityRef `json:"authoredBy,omitempty"`
	// A reference to the definition that this definition is a draft of, if this is a draft definition.
	DraftOf *DefinitionReference `json:"draftOf,omitempty"`
	// The list of drafts associated with this definition, if this is not a draft definition.
	Drafts  *[]DefinitionReference `json:"drafts,omitempty"`
	Metrics *[]BuildMetric         `json:"metrics,omitempty"`
	// The quality of the definition document (draft, etc.)
	Quality *DefinitionQuality `json:"quality,omitempty"`
	// The default queue for builds run against this definition.
	Queue *AgentPoolQueue `json:"queue,omitempty"`
}

// Represents a revision of a build definition.
type BuildDefinitionRevision struct {
	// The identity of the person or process that changed the definition.
	ChangedBy *webapi.IdentityRef `json:"changedBy,omitempty"`
	// The date and time that the definition was changed.
	ChangedDate *azuredevops.Time `json:"changedDate,omitempty"`
	// The change type (add, edit, delete).
	ChangeType *AuditAction `json:"changeType,omitempty"`
	// The comment associated with the change.
	Comment *string `json:"comment,omitempty"`
	// A link to the definition at this revision.
	DefinitionUrl *string `json:"definitionUrl,omitempty"`
	// The name of the definition.
	Name *string `json:"name,omitempty"`
	// The revision number.
	Revision *int `json:"revision,omitempty"`
}

type BuildDefinitionSourceProvider struct {
	// Uri of the associated definition
	DefinitionUri *string `json:"definitionUri,omitempty"`
	// fields associated with this build definition
	Fields *map[string]string `json:"fields,omitempty"`
	// Id of this source provider
	Id *int `json:"id,omitempty"`
	// The lst time this source provider was modified
	LastModified *azuredevops.Time `json:"lastModified,omitempty"`
	// Name of the source provider
	Name *string `json:"name,omitempty"`
	// Which trigger types are supported by this definition source provider
	SupportedTriggerTypes *DefinitionTriggerType `json:"supportedTriggerTypes,omitempty"`
}

// Represents a step in a build phase.
type BuildDefinitionStep struct {
	// Indicates whether this step should run even if a previous step fails.
	AlwaysRun *bool `json:"alwaysRun,omitempty"`
	// A condition that determines whether this step should run.
	Condition *string `json:"condition,omitempty"`
	// Indicates whether the phase should continue even if this step fails.
	ContinueOnError *bool `json:"continueOnError,omitempty"`
	// The display name for this step.
	DisplayName *string `json:"displayName,omitempty"`
	// Indicates whether the step is enabled.
	Enabled     *bool              `json:"enabled,omitempty"`
	Environment *map[string]string `json:"environment,omitempty"`
	Inputs      *map[string]string `json:"inputs,omitempty"`
	// The reference name for this step.
	RefName *string `json:"refName,omitempty"`
	// The task associated with this step.
	Task *TaskDefinitionReference `json:"task,omitempty"`
	// The time, in minutes, that this step is allowed to run.
	TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
}

// Represents a template from which new build definitions can be created.
type BuildDefinitionTemplate struct {
	// Indicates whether the template can be deleted.
	CanDelete *bool `json:"canDelete,omitempty"`
	// The template category.
	Category *string `json:"category,omitempty"`
	// An optional hosted agent queue for the template to use by default.
	DefaultHostedQueue *string `json:"defaultHostedQueue,omitempty"`
	// A description of the template.
	Description *string            `json:"description,omitempty"`
	Icons       *map[string]string `json:"icons,omitempty"`
	// The ID of the task whose icon is used when showing this template in the UI.
	IconTaskId *uuid.UUID `json:"iconTaskId,omitempty"`
	// The ID of the template.
	Id *string `json:"id,omitempty"`
	// The name of the template.
	Name *string `json:"name,omitempty"`
	// The actual template.
	Template *BuildDefinition `json:"template,omitempty"`
}

// For back-compat with extensions that use the old Steps format instead of Process and Phases
type BuildDefinitionTemplate3_2 struct {
	CanDelete          *bool               `json:"canDelete,omitempty"`
	Category           *string             `json:"category,omitempty"`
	DefaultHostedQueue *string             `json:"defaultHostedQueue,omitempty"`
	Description        *string             `json:"description,omitempty"`
	Icons              *map[string]string  `json:"icons,omitempty"`
	IconTaskId         *uuid.UUID          `json:"iconTaskId,omitempty"`
	Id                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Template           *BuildDefinition3_2 `json:"template,omitempty"`
}

// Represents a variable used by a build definition.
type BuildDefinitionVariable struct {
	// Indicates whether the value can be set at queue time.
	AllowOverride *bool `json:"allowOverride,omitempty"`
	// Indicates whether the variable's value is a secret.
	IsSecret *bool `json:"isSecret,omitempty"`
	// The value of the variable.
	Value *string `json:"value,omitempty"`
}

type BuildDeletedEvent struct {
	BuildId *int   `json:"buildId,omitempty"`
	Build   *Build `json:"build,omitempty"`
}

type BuildDeployment struct {
	Deployment  *BuildSummary       `json:"deployment,omitempty"`
	SourceBuild *XamlBuildReference `json:"sourceBuild,omitempty"`
}

type BuildEvent struct {
	Data       *[]string `json:"data,omitempty"`
	Identifier *string   `json:"identifier,omitempty"`
}

// Represents a build log.
type BuildLog struct {
	// The ID of the log.
	Id *int `json:"id,omitempty"`
	// The type of the log location.
	Type *string `json:"type,omitempty"`
	// A full link to the log resource.
	Url *string `json:"url,omitempty"`
	// The date and time the log was created.
	CreatedOn *azuredevops.Time `json:"createdOn,omitempty"`
	// The date and time the log was last changed.
	LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
	// The number of lines in the log.
	LineCount *uint64 `json:"lineCount,omitempty"`
}

// Represents a reference to a build log.
type BuildLogReference struct {
	// The ID of the log.
	Id *int `json:"id,omitempty"`
	// The type of the log location.
	Type *string `json:"type,omitempty"`
	// A full link to the log resource.
	Url *string `json:"url,omitempty"`
}

// Represents metadata about builds in the system.
type BuildMetric struct {
	// The date for the scope.
	Date *azuredevops.Time `json:"date,omitempty"`
	// The value.
	IntValue *int `json:"intValue,omitempty"`
	// The name of the metric.
	Name *string `json:"name,omitempty"`
	// The scope.
	Scope *string `json:"scope,omitempty"`
}

// Represents the application of an optional behavior to a build definition.
type BuildOption struct {
	// A reference to the build option.
	Definition *BuildOptionDefinitionReference `json:"definition,omitempty"`
	// Indicates whether the behavior is enabled.
	Enabled *bool              `json:"enabled,omitempty"`
	Inputs  *map[string]string `json:"inputs,omitempty"`
}

// Represents an optional behavior that can be applied to a build definition.
type BuildOptionDefinition struct {
	// The ID of the referenced build option.
	Id *uuid.UUID `json:"id,omitempty"`
	// The description.
	Description *string `json:"description,omitempty"`
	// The list of input groups defined for the build option.
	Groups *[]BuildOptionGroupDefinition `json:"groups,omitempty"`
	// The list of inputs defined for the build option.
	Inputs *[]BuildOptionInputDefinition `json:"inputs,omitempty"`
	// The name of the build option.
	Name *string `json:"name,omitempty"`
	// A value that indicates the relative order in which the behavior should be applied.
	Ordinal *int `json:"ordinal,omitempty"`
}

// Represents a reference to a build option definition.
type BuildOptionDefinitionReference struct {
	// The ID of the referenced build option.
	Id *uuid.UUID `json:"id,omitempty"`
}

// Represents a group of inputs for a build option.
type BuildOptionGroupDefinition struct {
	// The name of the group to display in the UI.
	DisplayName *string `json:"displayName,omitempty"`
	// Indicates whether the group is initially displayed as expanded in the UI.
	IsExpanded *bool `json:"isExpanded,omitempty"`
	// The internal name of the group.
	Name *string `json:"name,omitempty"`
}

// Represents an input for a build option.
type BuildOptionInputDefinition struct {
	// The default value.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// The name of the input group that this input belongs to.
	GroupName *string            `json:"groupName,omitempty"`
	Help      *map[string]string `json:"help,omitempty"`
	// The label for the input.
	Label *string `json:"label,omitempty"`
	// The name of the input.
	Name    *string            `json:"name,omitempty"`
	Options *map[string]string `json:"options,omitempty"`
	// Indicates whether the input is required to have a value.
	Required *bool `json:"required,omitempty"`
	// Indicates the type of the input value.
	Type *BuildOptionInputType `json:"type,omitempty"`
	// The rule that is applied to determine whether the input is visible in the UI.
	VisibleRule *string `json:"visibleRule,omitempty"`
}

type BuildOptionInputType string

type buildOptionInputTypeValuesType struct {
	String       BuildOptionInputType
	Boolean      BuildOptionInputType
	StringList   BuildOptionInputType
	Radio        BuildOptionInputType
	PickList     BuildOptionInputType
	MultiLine    BuildOptionInputType
	BranchFilter BuildOptionInputType
}

var BuildOptionInputTypeValues = buildOptionInputTypeValuesType{
	String:       "string",
	Boolean:      "boolean",
	StringList:   "stringList",
	Radio:        "radio",
	PickList:     "pickList",
	MultiLine:    "multiLine",
	BranchFilter: "branchFilter",
}

type BuildPhaseStatus string

type buildPhaseStatusValuesType struct {
	Unknown   BuildPhaseStatus
	Failed    BuildPhaseStatus
	Succeeded BuildPhaseStatus
}

var BuildPhaseStatusValues = buildPhaseStatusValuesType{
	// The state is not known.
	Unknown: "unknown",
	// The build phase completed unsuccessfully.
	Failed: "failed",
	// The build phase completed successfully.
	Succeeded: "succeeded",
}

// Represents resources used by a build process.
type BuildProcessResources struct {
	Endpoints      *[]ServiceEndpointReference `json:"endpoints,omitempty"`
	Files          *[]SecureFileReference      `json:"files,omitempty"`
	Queues         *[]AgentPoolQueueReference  `json:"queues,omitempty"`
	VariableGroups *[]VariableGroupReference   `json:"variableGroups,omitempty"`
}

type BuildProcessTemplate struct {
	Description      *string              `json:"description,omitempty"`
	FileExists       *bool                `json:"fileExists,omitempty"`
	Id               *int                 `json:"id,omitempty"`
	Parameters       *string              `json:"parameters,omitempty"`
	ServerPath       *string              `json:"serverPath,omitempty"`
	SupportedReasons *BuildReason         `json:"supportedReasons,omitempty"`
	TeamProject      *string              `json:"teamProject,omitempty"`
	TemplateType     *ProcessTemplateType `json:"templateType,omitempty"`
	Url              *string              `json:"url,omitempty"`
	Version          *string              `json:"version,omitempty"`
}

// Specifies the desired ordering of builds.
type BuildQueryOrder string

type buildQueryOrderValuesType struct {
	FinishTimeAscending  BuildQueryOrder
	FinishTimeDescending BuildQueryOrder
	QueueTimeDescending  BuildQueryOrder
	QueueTimeAscending   BuildQueryOrder
	StartTimeDescending  BuildQueryOrder
	StartTimeAscending   BuildQueryOrder
}

var BuildQueryOrderValues = buildQueryOrderValuesType{
	// Order by finish time ascending.
	FinishTimeAscending: "finishTimeAscending",
	// Order by finish time descending.
	FinishTimeDescending: "finishTimeDescending",
	// Order by queue time descending.
	QueueTimeDescending: "queueTimeDescending",
	// Order by queue time ascending.
	QueueTimeAscending: "queueTimeAscending",
	// Order by start time descending.
	StartTimeDescending: "startTimeDescending",
	// Order by start time ascending.
	StartTimeAscending: "startTimeAscending",
}

type BuildQueuedEvent struct {
	BuildId *int   `json:"buildId,omitempty"`
	Build   *Build `json:"build,omitempty"`
}

type BuildReason string

type buildReasonValuesType struct {
	None              BuildReason
	Manual            BuildReason
	IndividualCI      BuildReason
	BatchedCI         BuildReason
	Schedule          BuildReason
	ScheduleForced    BuildReason
	UserCreated       BuildReason
	ValidateShelveset BuildReason
	CheckInShelveset  BuildReason
	PullRequest       BuildReason
	BuildCompletion   BuildReason
	Triggered         BuildReason
	All               BuildReason
}

var BuildReasonValues = buildReasonValuesType{
	// No reason. This value should not be used.
	None: "none",
	// The build was started manually.
	Manual: "manual",
	// The build was started for the trigger TriggerType.ContinuousIntegration.
	IndividualCI: "individualCI",
	// The build was started for the trigger TriggerType.BatchedContinuousIntegration.
	BatchedCI: "batchedCI",
	// The build was started for the trigger TriggerType.Schedule.
	Schedule: "schedule",
	// The build was started for the trigger TriggerType.ScheduleForced.
	ScheduleForced: "scheduleForced",
	// The build was created by a user.
	UserCreated: "userCreated",
	// The build was started manually for private validation.
	ValidateShelveset: "validateShelveset",
	// The build was started for the trigger ContinuousIntegrationType.Gated.
	CheckInShelveset: "checkInShelveset",
	// The build was started by a pull request. Added in resource version 3.
	PullRequest: "pullRequest",
	// The build was started when another build completed.
	BuildCompletion: "buildCompletion",
	// The build was triggered for retention policy purposes.
	Triggered: "triggered",
	// All reasons.
	All: "all",
}

// Represents a reference to a build.
type BuildReference struct {
	Links interface{} `json:"_links,omitempty"`
	// The build number.
	BuildNumber *string `json:"buildNumber,omitempty"`
	// Indicates whether the build has been deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// The time that the build was completed.
	FinishTime *azuredevops.Time `json:"finishTime,omitempty"`
	// The ID of the build.
	Id *int `json:"id,omitempty"`
	// The time that the build was queued.
	QueueTime *azuredevops.Time `json:"queueTime,omitempty"`
	// The identity on whose behalf the build was queued.
	RequestedFor *webapi.IdentityRef `json:"requestedFor,omitempty"`
	// The build result.
	Result *BuildResult `json:"result,omitempty"`
	// The time that the build was started.
	StartTime *azuredevops.Time `json:"startTime,omitempty"`
	// The build status.
	Status *BuildStatus `json:"status,omitempty"`
}

// Represents information about a build report.
type BuildReportMetadata struct {
	// The Id of the build.
	BuildId *int `json:"buildId,omitempty"`
	// The content of the report.
	Content *string `json:"content,omitempty"`
	// The type of the report.
	Type *string `json:"type,omitempty"`
}

// Represents a repository used by a build definition.
type BuildRepository struct {
	// Indicates whether to checkout submodules.
	CheckoutSubmodules *bool `json:"checkoutSubmodules,omitempty"`
	// Indicates whether to clean the target folder when getting code from the repository.
	Clean *string `json:"clean,omitempty"`
	// The name of the default branch.
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	// The ID of the repository.
	Id *string `json:"id,omitempty"`
	// The friendly name of the repository.
	Name       *string            `json:"name,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
	// The root folder.
	RootFolder *string `json:"rootFolder,omitempty"`
	// The type of the repository.
	Type *string `json:"type,omitempty"`
	// The URL of the repository.
	Url *string `json:"url,omitempty"`
}

// Represents the result of validating a build request.
type BuildRequestValidationResult struct {
	// The message associated with the result.
	Message *string `json:"message,omitempty"`
	// The result.
	Result *ValidationResult `json:"result,omitempty"`
}

// Represents information about resources used by builds in the system.
type BuildResourceUsage struct {
	// The number of build agents.
	DistributedTaskAgents *int `json:"distributedTaskAgents,omitempty"`
	// The number of paid private agent slots.
	PaidPrivateAgentSlots *int `json:"paidPrivateAgentSlots,omitempty"`
	// The total usage.
	TotalUsage *int `json:"totalUsage,omitempty"`
	// The number of XAML controllers.
	XamlControllers *int `json:"xamlControllers,omitempty"`
}

// This is not a Flags enum because we don't want to set multiple statuses on a build. However, when adding values, please stick to powers of 2 as if it were a Flags enum This will ensure that things that key off multiple result types (like labelling sources) continue to work
type BuildResult string

type buildResultValuesType struct {
	None               BuildResult
	Succeeded          BuildResult
	PartiallySucceeded BuildResult
	Failed             BuildResult
	Canceled           BuildResult
}

var BuildResultValues = buildResultValuesType{
	// No result
	None: "none",
	// The build completed successfully.
	Succeeded: "succeeded",
	// The build completed compilation successfully but had other errors.
	PartiallySucceeded: "partiallySucceeded",
	// The build completed unsuccessfully.
	Failed: "failed",
	// The build was canceled before starting.
	Canceled: "canceled",
}

type BuildsDeletedEvent struct {
	BuildIds *[]int `json:"buildIds,omitempty"`
	// The ID of the definition.
	DefinitionId *int `json:"definitionId,omitempty"`
	// The ID of the project.
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
}

type BuildsDeletedEvent1 struct {
	BuildIds *[]int `json:"buildIds,omitempty"`
	// The ID of the definition.
	DefinitionId *int `json:"definitionId,omitempty"`
	// The ID of the project.
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
}

type BuildServer struct {
	Agents                    *[]BuildAgentReference        `json:"agents,omitempty"`
	Controller                *XamlBuildControllerReference `json:"controller,omitempty"`
	Id                        *int                          `json:"id,omitempty"`
	IsVirtual                 *bool                         `json:"isVirtual,omitempty"`
	MessageQueueUrl           *string                       `json:"messageQueueUrl,omitempty"`
	Name                      *string                       `json:"name,omitempty"`
	RequireClientCertificates *bool                         `json:"requireClientCertificates,omitempty"`
	Status                    *ServiceHostStatus            `json:"status,omitempty"`
	StatusChangedDate         *azuredevops.Time             `json:"statusChangedDate,omitempty"`
	Uri                       *string                       `json:"uri,omitempty"`
	Url                       *string                       `json:"url,omitempty"`
	Version                   *int                          `json:"version,omitempty"`
}

// Represents system-wide build settings.
type BuildSettings struct {
	// The number of days to keep records of deleted builds.
	DaysToKeepDeletedBuildsBeforeDestroy *int `json:"daysToKeepDeletedBuildsBeforeDestroy,omitempty"`
	// The default retention policy.
	DefaultRetentionPolicy *RetentionPolicy `json:"defaultRetentionPolicy,omitempty"`
	// The maximum retention policy.
	MaximumRetentionPolicy *RetentionPolicy `json:"maximumRetentionPolicy,omitempty"`
}

type BuildStatus string

type buildStatusValuesType struct {
	None       BuildStatus
	InProgress BuildStatus
	Completed  BuildStatus
	Cancelling BuildStatus
	Postponed  BuildStatus
	NotStarted BuildStatus
	All        BuildStatus
}

var BuildStatusValues = buildStatusValuesType{
	// No status.
	None: "none",
	// The build is currently in progress.
	InProgress: "inProgress",
	// The build has completed.
	Completed: "completed",
	// The build is cancelling
	Cancelling: "cancelling",
	// The build is inactive in the queue.
	Postponed: "postponed",
	// The build has not yet started.
	NotStarted: "notStarted",
	// All status.
	All: "all",
}

type BuildSummary struct {
	Build        *XamlBuildReference `json:"build,omitempty"`
	FinishTime   *azuredevops.Time   `json:"finishTime,omitempty"`
	KeepForever  *bool               `json:"keepForever,omitempty"`
	Quality      *string             `json:"quality,omitempty"`
	Reason       *BuildReason        `json:"reason,omitempty"`
	RequestedFor *webapi.IdentityRef `json:"requestedFor,omitempty"`
	StartTime    *azuredevops.Time   `json:"startTime,omitempty"`
	Status       *BuildStatus        `json:"status,omitempty"`
}

type BuildTagsAddedEvent struct {
	BuildId *int      `json:"buildId,omitempty"`
	Build   *Build    `json:"build,omitempty"`
	AllTags *[]string `json:"allTags,omitempty"`
	NewTags *[]string `json:"newTags,omitempty"`
}

type BuildUpdatedEvent struct {
	BuildId *int   `json:"buildId,omitempty"`
	Build   *Build `json:"build,omitempty"`
}

// Represents a workspace mapping.
type BuildWorkspace struct {
	Mappings *[]MappingDetails `json:"mappings,omitempty"`
}

// Represents a change associated with a build.
type Change struct {
	// The author of the change.
	Author *webapi.IdentityRef `json:"author,omitempty"`
	// The location of a user-friendly representation of the resource.
	DisplayUri *string `json:"displayUri,omitempty"`
	// The identifier for the change. For a commit, this would be the SHA1. For a TFVC changeset, this would be the changeset ID.
	Id *string `json:"id,omitempty"`
	// The location of the full representation of the resource.
	Location *string `json:"location,omitempty"`
	// The description of the change. This might be a commit message or changeset description.
	Message *string `json:"message,omitempty"`
	// Indicates whether the message was truncated.
	MessageTruncated *bool `json:"messageTruncated,omitempty"`
	// The person or process that pushed the change.
	Pusher *string `json:"pusher,omitempty"`
	// The timestamp for the change.
	Timestamp *azuredevops.Time `json:"timestamp,omitempty"`
	// The type of change. "commit", "changeset", etc.
	Type *string `json:"type,omitempty"`
}

type ConsoleLogEvent struct {
	BuildId          *int       `json:"buildId,omitempty"`
	Lines            *[]string  `json:"lines,omitempty"`
	StepRecordId     *uuid.UUID `json:"stepRecordId,omitempty"`
	TimelineId       *uuid.UUID `json:"timelineId,omitempty"`
	TimelineRecordId *uuid.UUID `json:"timelineRecordId,omitempty"`
}

type ContinuousDeploymentDefinition struct {
	// The connected service associated with the continuous deployment
	ConnectedService *core.WebApiConnectedServiceRef `json:"connectedService,omitempty"`
	// The definition associated with the continuous deployment
	Definition         *XamlDefinitionReference   `json:"definition,omitempty"`
	GitBranch          *string                    `json:"gitBranch,omitempty"`
	HostedServiceName  *string                    `json:"hostedServiceName,omitempty"`
	Project            *core.TeamProjectReference `json:"project,omitempty"`
	RepositoryId       *string                    `json:"repositoryId,omitempty"`
	StorageAccountName *string                    `json:"storageAccountName,omitempty"`
	SubscriptionId     *string                    `json:"subscriptionId,omitempty"`
	Website            *string                    `json:"website,omitempty"`
	Webspace           *string                    `json:"webspace,omitempty"`
}

// Represents a continuous integration (CI) trigger.
type ContinuousIntegrationTrigger struct {
	// Indicates whether changes should be batched while another CI build is running.
	BatchChanges  *bool     `json:"batchChanges,omitempty"`
	BranchFilters *[]string `json:"branchFilters,omitempty"`
	// The maximum number of simultaneous CI builds that will run per branch.
	MaxConcurrentBuildsPerBranch *int      `json:"maxConcurrentBuildsPerBranch,omitempty"`
	PathFilters                  *[]string `json:"pathFilters,omitempty"`
	// The polling interval, in seconds.
	PollingInterval *int `json:"pollingInterval,omitempty"`
	// The ID of the job used to poll an external repository.
	PollingJobId       *uuid.UUID `json:"pollingJobId,omitempty"`
	SettingsSourceType *int       `json:"settingsSourceType,omitempty"`
}

type ControllerStatus string

type controllerStatusValuesType struct {
	Unavailable ControllerStatus
	Available   ControllerStatus
	Offline     ControllerStatus
}

var ControllerStatusValues = controllerStatusValuesType{
	// Indicates that the build controller cannot be contacted.
	Unavailable: "unavailable",
	// Indicates that the build controller is currently available.
	Available: "available",
	// Indicates that the build controller has taken itself offline.
	Offline: "offline",
}

type DefinitionQuality string

type definitionQualityValuesType struct {
	Definition DefinitionQuality
	Draft      DefinitionQuality
}

var DefinitionQualityValues = definitionQualityValuesType{
	Definition: "definition",
	Draft:      "draft",
}

// Specifies the desired ordering of definitions.
type DefinitionQueryOrder string

type definitionQueryOrderValuesType struct {
	None                     DefinitionQueryOrder
	LastModifiedAscending    DefinitionQueryOrder
	LastModifiedDescending   DefinitionQueryOrder
	DefinitionNameAscending  DefinitionQueryOrder
	DefinitionNameDescending DefinitionQueryOrder
}

var DefinitionQueryOrderValues = definitionQueryOrderValuesType{
	// No order
	None: "none",
	// Order by created on/last modified time ascending.
	LastModifiedAscending: "lastModifiedAscending",
	// Order by created on/last modified time descending.
	LastModifiedDescending: "lastModifiedDescending",
	// Order by definition name ascending.
	DefinitionNameAscending: "definitionNameAscending",
	// Order by definition name descending.
	DefinitionNameDescending: "definitionNameDescending",
}

type DefinitionQueueStatus string

type definitionQueueStatusValuesType struct {
	Enabled  DefinitionQueueStatus
	Paused   DefinitionQueueStatus
	Disabled DefinitionQueueStatus
}

var DefinitionQueueStatusValues = definitionQueueStatusValuesType{
	// When enabled the definition queue allows builds to be queued by users, the system will queue scheduled, gated and continuous integration builds, and the queued builds will be started by the system.
	Enabled: "enabled",
	// When paused the definition queue allows builds to be queued by users and the system will queue scheduled, gated and continuous integration builds. Builds in the queue will not be started by the system.
	Paused: "paused",
	// When disabled the definition queue will not allow builds to be queued by users and the system will not queue scheduled, gated or continuous integration builds. Builds already in the queue will not be started by the system.
	Disabled: "disabled",
}

// Represents a reference to a definition.
type DefinitionReference struct {
	// The date this version of the definition was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The ID of the referenced definition.
	Id *int `json:"id,omitempty"`
	// The name of the referenced definition.
	Name *string `json:"name,omitempty"`
	// The folder path of the definition.
	Path *string `json:"path,omitempty"`
	// A reference to the project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// A value that indicates whether builds can be queued against this definition.
	QueueStatus *DefinitionQueueStatus `json:"queueStatus,omitempty"`
	// The definition revision number.
	Revision *int `json:"revision,omitempty"`
	// The type of the definition.
	Type *DefinitionType `json:"type,omitempty"`
	// The definition's URI.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the definition.
	Url *string `json:"url,omitempty"`
}

type DefinitionResourceReference struct {
	// Indicates whether the resource is authorized for use.
	Authorized *bool `json:"authorized,omitempty"`
	// The id of the resource.
	Id *string `json:"id,omitempty"`
	// A friendly name for the resource.
	Name *string `json:"name,omitempty"`
	// The type of the resource.
	Type *string `json:"type,omitempty"`
}

type DefinitionTriggerType string

type definitionTriggerTypeValuesType struct {
	None                         DefinitionTriggerType
	ContinuousIntegration        DefinitionTriggerType
	BatchedContinuousIntegration DefinitionTriggerType
	Schedule                     DefinitionTriggerType
	GatedCheckIn                 DefinitionTriggerType
	BatchedGatedCheckIn          DefinitionTriggerType
	PullRequest                  DefinitionTriggerType
	BuildCompletion              DefinitionTriggerType
	All                          DefinitionTriggerType
}

var DefinitionTriggerTypeValues = definitionTriggerTypeValuesType{
	// Manual builds only.
	None: "none",
	// A build should be started for each changeset.
	ContinuousIntegration: "continuousIntegration",
	// A build should be started for multiple changesets at a time at a specified interval.
	BatchedContinuousIntegration: "batchedContinuousIntegration",
	// A build should be started on a specified schedule whether or not changesets exist.
	Schedule: "schedule",
	// A validation build should be started for each check-in.
	GatedCheckIn: "gatedCheckIn",
	// A validation build should be started for each batch of check-ins.
	BatchedGatedCheckIn: "batchedGatedCheckIn",
	// A build should be triggered when a GitHub pull request is created or updated. Added in resource version 3
	PullRequest: "pullRequest",
	// A build should be triggered when another build completes.
	BuildCompletion: "buildCompletion",
	// All types.
	All: "all",
}

type DefinitionType string

type definitionTypeValuesType struct {
	Xaml  DefinitionType
	Build DefinitionType
}

var DefinitionTypeValues = definitionTypeValuesType{
	Xaml:  "xaml",
	Build: "build",
}

type DeleteOptions string

type deleteOptionsValuesType struct {
	None         DeleteOptions
	DropLocation DeleteOptions
	TestResults  DeleteOptions
	Label        DeleteOptions
	Details      DeleteOptions
	Symbols      DeleteOptions
	All          DeleteOptions
}

var DeleteOptionsValues = deleteOptionsValuesType{
	// No data should be deleted. This value should not be used.
	None: "none",
	// The drop location should be deleted.
	DropLocation: "dropLocation",
	// The test results should be deleted.
	TestResults: "testResults",
	// The version control label should be deleted.
	Label: "label",
	// The build should be deleted.
	Details: "details",
	// Published symbols should be deleted.
	Symbols: "symbols",
	// All data should be deleted.
	All: "all",
}

// Represents a dependency.
type Dependency struct {
	// The event. The dependency is satisfied when the referenced object emits this event.
	Event *string `json:"event,omitempty"`
	// The scope. This names the object referenced by the dependency.
	Scope *string `json:"scope,omitempty"`
}

// Represents the data from the build information nodes for type "DeploymentInformation" for xaml builds
type Deployment struct {
	Type *string `json:"type,omitempty"`
}

// Deployment information for type "Build"
type DeploymentBuild struct {
	Type    *string `json:"type,omitempty"`
	BuildId *int    `json:"buildId,omitempty"`
}

// Deployment information for type "Deploy"
type DeploymentDeploy struct {
	Type    *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Deployment information for type "Test"
type DeploymentTest struct {
	Type  *string `json:"type,omitempty"`
	RunId *int    `json:"runId,omitempty"`
}

// Represents a build process supported by the build definition designer.
type DesignerProcess struct {
	Phases *[]Phase `json:"phases,omitempty"`
	// The target for the build process.
	Target *DesignerProcessTarget `json:"target,omitempty"`
}

// Represents the target for the build process.
type DesignerProcessTarget struct {
	// Agent specification for the build process.
	AgentSpecification *AgentSpecification `json:"agentSpecification,omitempty"`
}

type DockerProcess struct {
	Target *DockerProcessTarget `json:"target,omitempty"`
}

// Represents the target for the docker build process.
type DockerProcessTarget struct {
	// Agent specification for the build process.
	AgentSpecification *AgentSpecification `json:"agentSpecification,omitempty"`
}

// Represents a folder that contains build definitions.
type Folder struct {
	// The process or person who created the folder.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// The date the folder was created.
	CreatedOn *azuredevops.Time `json:"createdOn,omitempty"`
	// The description.
	Description *string `json:"description,omitempty"`
	// The process or person that last changed the folder.
	LastChangedBy *webapi.IdentityRef `json:"lastChangedBy,omitempty"`
	// The date the folder was last changed.
	LastChangedDate *azuredevops.Time `json:"lastChangedDate,omitempty"`
	// The full path.
	Path *string `json:"path,omitempty"`
	// The project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
}

// Specifies the desired ordering of folders.
type FolderQueryOrder string

type folderQueryOrderValuesType struct {
	None             FolderQueryOrder
	FolderAscending  FolderQueryOrder
	FolderDescending FolderQueryOrder
}

var FolderQueryOrderValues = folderQueryOrderValuesType{
	// No order
	None: "none",
	// Order by folder name and path ascending.
	FolderAscending: "folderAscending",
	// Order by folder name and path descending.
	FolderDescending: "folderDescending",
}

// Represents the ability to build forks of the selected repository.
type Forks struct {
	// Indicates whether a build should use secrets when building forks of the selected repository.
	AllowSecrets *bool `json:"allowSecrets,omitempty"`
	// Indicates whether the trigger should queue builds for forks of the selected repository.
	Enabled *bool `json:"enabled,omitempty"`
}

// Represents a gated check-in trigger.
type GatedCheckInTrigger struct {
	PathFilters *[]string `json:"pathFilters,omitempty"`
	// Indicates whether CI triggers should run after the gated check-in succeeds.
	RunContinuousIntegration *bool `json:"runContinuousIntegration,omitempty"`
	// Indicates whether to take workspace mappings into account when determining whether a build should run.
	UseWorkspaceMappings *bool `json:"useWorkspaceMappings,omitempty"`
}

type GetOption string

type getOptionValuesType struct {
	LatestOnQueue GetOption
	LatestOnBuild GetOption
	Custom        GetOption
}

var GetOptionValues = getOptionValuesType{
	// Use the latest changeset at the time the build is queued.
	LatestOnQueue: "latestOnQueue",
	// Use the latest changeset at the time the build is started.
	LatestOnBuild: "latestOnBuild",
	// A user-specified version has been supplied.
	Custom: "custom",
}

// Data representation of an information node associated with a build
type InformationNode struct {
	// Fields of the information node
	Fields *map[string]string `json:"fields,omitempty"`
	// Process or person that last modified this node
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// Date this node was last modified
	LastModifiedDate *azuredevops.Time `json:"lastModifiedDate,omitempty"`
	// Node Id of this information node
	NodeId *int `json:"nodeId,omitempty"`
	// Id of parent node (xml tree)
	ParentId *int `json:"parentId,omitempty"`
	// The type of the information node
	Type *string `json:"type,omitempty"`
}

// Represents an issue (error, warning) associated with a build.
type Issue struct {
	// The category.
	Category *string            `json:"category,omitempty"`
	Data     *map[string]string `json:"data,omitempty"`
	// A description of the issue.
	Message *string `json:"message,omitempty"`
	// The type (error, warning) of the issue.
	Type *IssueType `json:"type,omitempty"`
}

type IssueType string

type issueTypeValuesType struct {
	Error   IssueType
	Warning IssueType
}

var IssueTypeValues = issueTypeValuesType{
	Error:   "error",
	Warning: "warning",
}

type JustInTimeProcess struct {
}

// Represents an entry in a workspace mapping.
type MappingDetails struct {
	// The local path.
	LocalPath *string `json:"localPath,omitempty"`
	// The mapping type.
	MappingType *string `json:"mappingType,omitempty"`
	// The server path.
	ServerPath *string `json:"serverPath,omitempty"`
}

// Represents options for running a phase against multiple agents.
type MultipleAgentExecutionOptions struct {
	// Indicates the type of execution options.
	Type *int `json:"type,omitempty"`
	// Indicates whether failure on one agent should prevent the phase from running on other agents.
	ContinueOnError *bool `json:"continueOnError,omitempty"`
	// The maximum number of agents to use simultaneously.
	MaxConcurrency *int `json:"maxConcurrency,omitempty"`
}

// Represents a phase of a build definition.
type Phase struct {
	// The condition that must be true for this phase to execute.
	Condition    *string       `json:"condition,omitempty"`
	Dependencies *[]Dependency `json:"dependencies,omitempty"`
	// The job authorization scope for builds queued against this definition.
	JobAuthorizationScope *BuildAuthorizationScope `json:"jobAuthorizationScope,omitempty"`
	// The cancellation timeout, in minutes, for builds queued against this definition.
	JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
	// The job execution timeout, in minutes, for builds queued against this definition.
	JobTimeoutInMinutes *int `json:"jobTimeoutInMinutes,omitempty"`
	// The name of the phase.
	Name *string `json:"name,omitempty"`
	// The unique ref name of the phase.
	RefName *string                `json:"refName,omitempty"`
	Steps   *[]BuildDefinitionStep `json:"steps,omitempty"`
	// The target (agent, server, etc.) for this phase.
	Target    *PhaseTarget                        `json:"target,omitempty"`
	Variables *map[string]BuildDefinitionVariable `json:"variables,omitempty"`
}

// Represents the target of a phase.
type PhaseTarget struct {
	// The type of the target.
	Type *int `json:"type,omitempty"`
}

type ProcessTemplateType string

type processTemplateTypeValuesType struct {
	Custom  ProcessTemplateType
	Default ProcessTemplateType
	Upgrade ProcessTemplateType
}

var ProcessTemplateTypeValues = processTemplateTypeValuesType{
	// Indicates a custom template.
	Custom: "custom",
	// Indicates a default template.
	Default: "default",
	// Indicates an upgrade template.
	Upgrade: "upgrade",
}

// Represents a pull request object.  These are retrieved from Source Providers.
type PullRequest struct {
	// The links to other objects related to this object.
	Links interface{} `json:"_links,omitempty"`
	// Author of the pull request.
	Author *webapi.IdentityRef `json:"author,omitempty"`
	// Current state of the pull request, e.g. open, merged, closed, conflicts, etc.
	CurrentState *string `json:"currentState,omitempty"`
	// Description for the pull request.
	Description *string `json:"description,omitempty"`
	// Unique identifier for the pull request
	Id *string `json:"id,omitempty"`
	// The name of the provider this pull request is associated with.
	ProviderName *string `json:"providerName,omitempty"`
	// Source branch ref of this pull request
	SourceBranchRef *string `json:"sourceBranchRef,omitempty"`
	// Owner of the source repository of this pull request
	SourceRepositoryOwner *string `json:"sourceRepositoryOwner,omitempty"`
	// Target branch ref of this pull request
	TargetBranchRef *string `json:"targetBranchRef,omitempty"`
	// Owner of the target repository of this pull request
	TargetRepositoryOwner *string `json:"targetRepositoryOwner,omitempty"`
	// Title of the pull request.
	Title *string `json:"title,omitempty"`
}

// Represents a pull request trigger.
type PullRequestTrigger struct {
	// Indicates if an update to a PR should delete current in-progress builds.
	AutoCancel                           *bool     `json:"autoCancel,omitempty"`
	BranchFilters                        *[]string `json:"branchFilters,omitempty"`
	Forks                                *Forks    `json:"forks,omitempty"`
	IsCommentRequiredForPullRequest      *bool     `json:"isCommentRequiredForPullRequest,omitempty"`
	PathFilters                          *[]string `json:"pathFilters,omitempty"`
	RequireCommentsForNonTeamMembersOnly *bool     `json:"requireCommentsForNonTeamMembersOnly,omitempty"`
	SettingsSourceType                   *int      `json:"settingsSourceType,omitempty"`
}

type QueryDeletedOption string

type queryDeletedOptionValuesType struct {
	ExcludeDeleted QueryDeletedOption
	IncludeDeleted QueryDeletedOption
	OnlyDeleted    QueryDeletedOption
}

var QueryDeletedOptionValues = queryDeletedOptionValuesType{
	// Include only non-deleted builds.
	ExcludeDeleted: "excludeDeleted",
	// Include deleted and non-deleted builds.
	IncludeDeleted: "includeDeleted",
	// Include only deleted builds.
	OnlyDeleted: "onlyDeleted",
}

// [Flags]
type QueueOptions string

type queueOptionsValuesType struct {
	None     QueueOptions
	DoNotRun QueueOptions
}

var QueueOptionsValues = queueOptionsValuesType{
	// No queue options
	None: "none",
	// Create a plan Id for the build, do not run it
	DoNotRun: "doNotRun",
}

type QueuePriority string

type queuePriorityValuesType struct {
	Low         QueuePriority
	BelowNormal QueuePriority
	Normal      QueuePriority
	AboveNormal QueuePriority
	High        QueuePriority
}

var QueuePriorityValues = queuePriorityValuesType{
	// Low priority.
	Low: "low",
	// Below normal priority.
	BelowNormal: "belowNormal",
	// Normal priority.
	Normal: "normal",
	// Above normal priority.
	AboveNormal: "aboveNormal",
	// High priority.
	High: "high",
}

type RealtimeBuildEvent struct {
	BuildId *int `json:"buildId,omitempty"`
}

type RepositoryCleanOptions string

type repositoryCleanOptionsValuesType struct {
	Source             RepositoryCleanOptions
	SourceAndOutputDir RepositoryCleanOptions
	SourceDir          RepositoryCleanOptions
	AllBuildDir        RepositoryCleanOptions
}

var RepositoryCleanOptionsValues = repositoryCleanOptionsValuesType{
	Source:             "source",
	SourceAndOutputDir: "sourceAndOutputDir",
	// Re-create $(build.sourcesDirectory)
	SourceDir: "sourceDir",
	// Re-create $(agnet.buildDirectory) which contains $(build.sourcesDirectory), $(build.binariesDirectory) and any folders that left from previous build.
	AllBuildDir: "allBuildDir",
}

// Represents a repository's webhook returned from a source provider.
type RepositoryWebhook struct {
	// The friendly name of the repository.
	Name  *string                  `json:"name,omitempty"`
	Types *[]DefinitionTriggerType `json:"types,omitempty"`
	// The URL of the repository.
	Url *string `json:"url,omitempty"`
}

// Represents a reference to a resource.
type ResourceReference struct {
	// An alias to be used when referencing the resource.
	Alias *string `json:"alias,omitempty"`
}

type ResultSet string

type resultSetValuesType struct {
	All ResultSet
	Top ResultSet
}

var ResultSetValues = resultSetValuesType{
	// Include all repositories
	All: "all",
	// Include most relevant repositories for user
	Top: "top",
}

// Represents a retention policy for a build definition.
type RetentionPolicy struct {
	Artifacts             *[]string `json:"artifacts,omitempty"`
	ArtifactTypesToDelete *[]string `json:"artifactTypesToDelete,omitempty"`
	Branches              *[]string `json:"branches,omitempty"`
	// The number of days to keep builds.
	DaysToKeep *int `json:"daysToKeep,omitempty"`
	// Indicates whether the build record itself should be deleted.
	DeleteBuildRecord *bool `json:"deleteBuildRecord,omitempty"`
	// Indicates whether to delete test results associated with the build.
	DeleteTestResults *bool `json:"deleteTestResults,omitempty"`
	// The minimum number of builds to keep.
	MinimumToKeep *int `json:"minimumToKeep,omitempty"`
}

type Schedule struct {
	BranchFilters *[]string `json:"branchFilters,omitempty"`
	// Days for a build (flags enum for days of the week)
	DaysToBuild *ScheduleDays `json:"daysToBuild,omitempty"`
	// The Job Id of the Scheduled job that will queue the scheduled build. Since a single trigger can have multiple schedules and we want a single job to process a single schedule (since each schedule has a list of branches to build), the schedule itself needs to define the Job Id. This value will be filled in when a definition is added or updated.  The UI does not provide it or use it.
	ScheduleJobId *uuid.UUID `json:"scheduleJobId,omitempty"`
	// Flag to determine if this schedule should only build if the associated source has been changed.
	ScheduleOnlyWithChanges *bool `json:"scheduleOnlyWithChanges,omitempty"`
	// Local timezone hour to start
	StartHours *int `json:"startHours,omitempty"`
	// Local timezone minute to start
	StartMinutes *int `json:"startMinutes,omitempty"`
	// Time zone of the build schedule (String representation of the time zone ID)
	TimeZoneId *string `json:"timeZoneId,omitempty"`
}

type ScheduleDays string

type scheduleDaysValuesType struct {
	None      ScheduleDays
	Monday    ScheduleDays
	Tuesday   ScheduleDays
	Wednesday ScheduleDays
	Thursday  ScheduleDays
	Friday    ScheduleDays
	Saturday  ScheduleDays
	Sunday    ScheduleDays
	All       ScheduleDays
}

var ScheduleDaysValues = scheduleDaysValuesType{
	// Do not run.
	None: "none",
	// Run on Monday.
	Monday: "monday",
	// Run on Tuesday.
	Tuesday: "tuesday",
	// Run on Wednesday.
	Wednesday: "wednesday",
	// Run on Thursday.
	Thursday: "thursday",
	// Run on Friday.
	Friday: "friday",
	// Run on Saturday.
	Saturday: "saturday",
	// Run on Sunday.
	Sunday: "sunday",
	// Run on all days of the week.
	All: "all",
}

// Represents a schedule trigger.
type ScheduleTrigger struct {
	Schedules *[]Schedule `json:"schedules,omitempty"`
}

// Represents a reference to a secure file.
type SecureFileReference struct {
	// An alias to be used when referencing the resource.
	Alias *string `json:"alias,omitempty"`
	// The ID of the secure file.
	Id *uuid.UUID `json:"id,omitempty"`
}

// Represents a phase target that runs on the server.
type ServerTarget struct {
	// The type of the target.
	Type *int `json:"type,omitempty"`
	// The execution options.
	ExecutionOptions *ServerTargetExecutionOptions `json:"executionOptions,omitempty"`
}

// Represents options for running a phase on the server.
type ServerTargetExecutionOptions struct {
	// The type.
	Type *int `json:"type,omitempty"`
}

// Represents a referenec to a service endpoint.
type ServiceEndpointReference struct {
	// An alias to be used when referencing the resource.
	Alias *string `json:"alias,omitempty"`
	// The ID of the service endpoint.
	Id *uuid.UUID `json:"id,omitempty"`
}

type ServiceHostStatus string

type serviceHostStatusValuesType struct {
	Online  ServiceHostStatus
	Offline ServiceHostStatus
}

var ServiceHostStatusValues = serviceHostStatusValuesType{
	// The service host is currently connected and accepting commands.
	Online: "online",
	// The service host is currently disconnected and not accepting commands.
	Offline: "offline",
}

type SourceProviderAttributes struct {
	// The name of the source provider.
	Name *string `json:"name,omitempty"`
	// The capabilities supported by this source provider.
	SupportedCapabilities *map[string]bool `json:"supportedCapabilities,omitempty"`
	// The types of triggers supported by this source provider.
	SupportedTriggers *[]SupportedTrigger `json:"supportedTriggers,omitempty"`
}

type SourceProviderAvailability string

type sourceProviderAvailabilityValuesType struct {
	Hosted     SourceProviderAvailability
	OnPremises SourceProviderAvailability
	All        SourceProviderAvailability
}

var SourceProviderAvailabilityValues = sourceProviderAvailabilityValuesType{
	// The source provider is available in the hosted environment.
	Hosted: "hosted",
	// The source provider is available in the on-premises environment.
	OnPremises: "onPremises",
	// The source provider is available in all environments.
	All: "all",
}

// Represents a work item related to some source item. These are retrieved from Source Providers.
type SourceRelatedWorkItem struct {
	Links interface{} `json:"_links,omitempty"`
	// Identity ref for the person that the work item is assigned to.
	AssignedTo *webapi.IdentityRef `json:"assignedTo,omitempty"`
	// Current state of the work item, e.g. Active, Resolved, Closed, etc.
	CurrentState *string `json:"currentState,omitempty"`
	// Long description for the work item.
	Description *string `json:"description,omitempty"`
	// Unique identifier for the work item
	Id *string `json:"id,omitempty"`
	// The name of the provider the work item is associated with.
	ProviderName *string `json:"providerName,omitempty"`
	// Short name for the work item.
	Title *string `json:"title,omitempty"`
	// Type of work item, e.g. Bug, Task, User Story, etc.
	Type *string `json:"type,omitempty"`
}

// A set of repositories returned from the source provider.
type SourceRepositories struct {
	// A token used to continue this paged request; 'null' if the request is complete
	ContinuationToken *string `json:"continuationToken,omitempty"`
	// The number of repositories requested for each page
	PageLength *int `json:"pageLength,omitempty"`
	// A list of repositories
	Repositories *[]SourceRepository `json:"repositories,omitempty"`
	// The total number of pages, or '-1' if unknown
	TotalPageCount *int `json:"totalPageCount,omitempty"`
}

// Represents a repository returned from a source provider.
type SourceRepository struct {
	// The name of the default branch.
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	// The full name of the repository.
	FullName *string `json:"fullName,omitempty"`
	// The ID of the repository.
	Id *string `json:"id,omitempty"`
	// The friendly name of the repository.
	Name       *string            `json:"name,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
	// The name of the source provider the repository is from.
	SourceProviderName *string `json:"sourceProviderName,omitempty"`
	// The URL of the repository.
	Url *string `json:"url,omitempty"`
}

// Represents an item in a repository from a source provider.
type SourceRepositoryItem struct {
	// Whether the item is able to have sub-items (e.g., is a folder).
	IsContainer *bool `json:"isContainer,omitempty"`
	// The full path of the item, relative to the root of the repository.
	Path *string `json:"path,omitempty"`
	// The type of the item (folder, file, etc).
	Type *string `json:"type,omitempty"`
	// The URL of the item.
	Url *string `json:"url,omitempty"`
}

type SupportedTrigger struct {
	// The default interval to wait between polls (only relevant when NotificationType is Polling).
	DefaultPollingInterval *int `json:"defaultPollingInterval,omitempty"`
	// How the trigger is notified of changes.
	NotificationType *string `json:"notificationType,omitempty"`
	// The capabilities supported by this trigger.
	SupportedCapabilities *map[string]SupportLevel `json:"supportedCapabilities,omitempty"`
	// The type of trigger.
	Type *DefinitionTriggerType `json:"type,omitempty"`
}

type SupportLevel string

type supportLevelValuesType struct {
	Unsupported SupportLevel
	Supported   SupportLevel
	Required    SupportLevel
}

var SupportLevelValues = supportLevelValuesType{
	// The functionality is not supported.
	Unsupported: "unsupported",
	// The functionality is supported.
	Supported: "supported",
	// The functionality is required.
	Required: "required",
}

// Represents a Subversion mapping entry.
type SvnMappingDetails struct {
	// The depth.
	Depth *int `json:"depth,omitempty"`
	// Indicates whether to ignore externals.
	IgnoreExternals *bool `json:"ignoreExternals,omitempty"`
	// The local path.
	LocalPath *string `json:"localPath,omitempty"`
	// The revision.
	Revision *string `json:"revision,omitempty"`
	// The server path.
	ServerPath *string `json:"serverPath,omitempty"`
}

// Represents a subversion workspace.
type SvnWorkspace struct {
	Mappings *[]SvnMappingDetails `json:"mappings,omitempty"`
}

// Represents a reference to an agent pool.
type TaskAgentPoolReference struct {
	// The pool ID.
	Id *int `json:"id,omitempty"`
	// A value indicating whether or not this pool is managed by the service.
	IsHosted *bool `json:"isHosted,omitempty"`
	// The pool name.
	Name *string `json:"name,omitempty"`
}

// A reference to a task definition.
type TaskDefinitionReference struct {
	// The type of task (task or task group).
	DefinitionType *string `json:"definitionType,omitempty"`
	// The ID of the task.
	Id *uuid.UUID `json:"id,omitempty"`
	// The version of the task.
	VersionSpec *string `json:"versionSpec,omitempty"`
}

// Represents a reference to a plan group.
type TaskOrchestrationPlanGroupReference struct {
	// The name of the plan group.
	PlanGroup *string `json:"planGroup,omitempty"`
	// The project ID.
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
}

type TaskOrchestrationPlanGroupsStartedEvent struct {
	PlanGroups *[]TaskOrchestrationPlanGroupReference `json:"planGroups,omitempty"`
}

// Represents a reference to an orchestration plan.
type TaskOrchestrationPlanReference struct {
	// The type of the plan.
	OrchestrationType *int `json:"orchestrationType,omitempty"`
	// The ID of the plan.
	PlanId *uuid.UUID `json:"planId,omitempty"`
}

// Represents a reference to a task.
type TaskReference struct {
	// The ID of the task definition.
	Id *uuid.UUID `json:"id,omitempty"`
	// The name of the task definition.
	Name *string `json:"name,omitempty"`
	// The version of the task definition.
	Version *string `json:"version,omitempty"`
}

type TaskResult string

type taskResultValuesType struct {
	Succeeded           TaskResult
	SucceededWithIssues TaskResult
	Failed              TaskResult
	Canceled            TaskResult
	Skipped             TaskResult
	Abandoned           TaskResult
}

var TaskResultValues = taskResultValuesType{
	Succeeded:           "succeeded",
	SucceededWithIssues: "succeededWithIssues",
	Failed:              "failed",
	Canceled:            "canceled",
	Skipped:             "skipped",
	Abandoned:           "abandoned",
}

// Represents the timeline of a build.
type Timeline struct {
	// The change ID.
	ChangeId *int `json:"changeId,omitempty"`
	// The ID of the timeline.
	Id *uuid.UUID `json:"id,omitempty"`
	// The REST URL of the timeline.
	Url *string `json:"url,omitempty"`
	// The process or person that last changed the timeline.
	LastChangedBy *uuid.UUID `json:"lastChangedBy,omitempty"`
	// The time the timeline was last changed.
	LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
	Records       *[]TimelineRecord `json:"records,omitempty"`
}

type TimelineAttempt struct {
	// Gets or sets the attempt of the record.
	Attempt *int `json:"attempt,omitempty"`
	// Gets or sets the record identifier located within the specified timeline.
	RecordId *uuid.UUID `json:"recordId,omitempty"`
	// Gets or sets the timeline identifier which owns the record representing this attempt.
	TimelineId *uuid.UUID `json:"timelineId,omitempty"`
}

// Represents an entry in a build's timeline.
type TimelineRecord struct {
	Links interface{} `json:"_links,omitempty"`
	// Attempt number of record.
	Attempt *int `json:"attempt,omitempty"`
	// The change ID.
	ChangeId *int `json:"changeId,omitempty"`
	// A string that indicates the current operation.
	CurrentOperation *string `json:"currentOperation,omitempty"`
	// A reference to a sub-timeline.
	Details *TimelineReference `json:"details,omitempty"`
	// The number of errors produced by this operation.
	ErrorCount *int `json:"errorCount,omitempty"`
	// The finish time.
	FinishTime *azuredevops.Time `json:"finishTime,omitempty"`
	// The ID of the record.
	Id *uuid.UUID `json:"id,omitempty"`
	// String identifier that is consistent across attempts.
	Identifier *string  `json:"identifier,omitempty"`
	Issues     *[]Issue `json:"issues,omitempty"`
	// The time the record was last modified.
	LastModified *azuredevops.Time `json:"lastModified,omitempty"`
	// A reference to the log produced by this operation.
	Log *BuildLogReference `json:"log,omitempty"`
	// The name.
	Name *string `json:"name,omitempty"`
	// An ordinal value relative to other records.
	Order *int `json:"order,omitempty"`
	// The ID of the record's parent.
	ParentId *uuid.UUID `json:"parentId,omitempty"`
	// The current completion percentage.
	PercentComplete  *int               `json:"percentComplete,omitempty"`
	PreviousAttempts *[]TimelineAttempt `json:"previousAttempts,omitempty"`
	// The result.
	Result *TaskResult `json:"result,omitempty"`
	// The result code.
	ResultCode *string `json:"resultCode,omitempty"`
	// The start time.
	StartTime *azuredevops.Time `json:"startTime,omitempty"`
	// The state of the record.
	State *TimelineRecordState `json:"state,omitempty"`
	// A reference to the task represented by this timeline record.
	Task *TaskReference `json:"task,omitempty"`
	// The type of the record.
	Type *string `json:"type,omitempty"`
	// The REST URL of the timeline record.
	Url *string `json:"url,omitempty"`
	// The number of warnings produced by this operation.
	WarningCount *int `json:"warningCount,omitempty"`
	// The name of the agent running the operation.
	WorkerName *string `json:"workerName,omitempty"`
}

type TimelineRecordState string

type timelineRecordStateValuesType struct {
	Pending    TimelineRecordState
	InProgress TimelineRecordState
	Completed  TimelineRecordState
}

var TimelineRecordStateValues = timelineRecordStateValuesType{
	Pending:    "pending",
	InProgress: "inProgress",
	Completed:  "completed",
}

type TimelineRecordsUpdatedEvent struct {
	BuildId         *int              `json:"buildId,omitempty"`
	TimelineRecords *[]TimelineRecord `json:"timelineRecords,omitempty"`
}

// Represents a reference to a timeline.
type TimelineReference struct {
	// The change ID.
	ChangeId *int `json:"changeId,omitempty"`
	// The ID of the timeline.
	Id *uuid.UUID `json:"id,omitempty"`
	// The REST URL of the timeline.
	Url *string `json:"url,omitempty"`
}

type ValidationResult string

type validationResultValuesType struct {
	Ok      ValidationResult
	Warning ValidationResult
	Error   ValidationResult
}

var ValidationResultValues = validationResultValuesType{
	Ok:      "ok",
	Warning: "warning",
	Error:   "error",
}

// Represents a variable group.
type VariableGroup struct {
	// The Name of the variable group.
	Alias *string `json:"alias,omitempty"`
	// The ID of the variable group.
	Id *int `json:"id,omitempty"`
	// The description.
	Description *string `json:"description,omitempty"`
	// The name of the variable group.
	Name *string `json:"name,omitempty"`
	// The type of the variable group.
	Type      *string                             `json:"type,omitempty"`
	Variables *map[string]BuildDefinitionVariable `json:"variables,omitempty"`
}

// Represents a reference to a variable group.
type VariableGroupReference struct {
	// The Name of the variable group.
	Alias *string `json:"alias,omitempty"`
	// The ID of the variable group.
	Id *int `json:"id,omitempty"`
}

// Represents options for running a phase based on values specified by a list of variables.
type VariableMultipliersAgentExecutionOptions struct {
	// Indicates the type of execution options.
	Type *int `json:"type,omitempty"`
	// Indicates whether failure on one agent should prevent the phase from running on other agents.
	ContinueOnError *bool `json:"continueOnError,omitempty"`
	// The maximum number of agents to use in parallel.
	MaxConcurrency *int      `json:"maxConcurrency,omitempty"`
	Multipliers    *[]string `json:"multipliers,omitempty"`
}

// Represents options for running a phase based on values specified by a list of variables.
type VariableMultipliersServerExecutionOptions struct {
	// The type.
	Type *int `json:"type,omitempty"`
	// Indicates whether failure of one job should prevent the phase from running in other jobs.
	ContinueOnError *bool `json:"continueOnError,omitempty"`
	// The maximum number of server jobs to run in parallel.
	MaxConcurrency *int      `json:"maxConcurrency,omitempty"`
	Multipliers    *[]string `json:"multipliers,omitempty"`
}

// Mapping for a workspace
type WorkspaceMapping struct {
	// Uri of the associated definition
	DefinitionUri *string `json:"definitionUri,omitempty"`
	// Depth of this mapping
	Depth *int `json:"depth,omitempty"`
	// local location of the definition
	LocalItem *string `json:"localItem,omitempty"`
	// type of workspace mapping
	MappingType *WorkspaceMappingType `json:"mappingType,omitempty"`
	// Server location of the definition
	ServerItem *string `json:"serverItem,omitempty"`
	// Id of the workspace
	WorkspaceId *int `json:"workspaceId,omitempty"`
}

type WorkspaceMappingType string

type workspaceMappingTypeValuesType struct {
	Map   WorkspaceMappingType
	Cloak WorkspaceMappingType
}

var WorkspaceMappingTypeValues = workspaceMappingTypeValuesType{
	// The path is mapped in the workspace.
	Map: "map",
	// The path is cloaked in the workspace.
	Cloak: "cloak",
}

type WorkspaceTemplate struct {
	// Uri of the associated definition
	DefinitionUri *string `json:"definitionUri,omitempty"`
	// The identity that last modified this template
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// The last time this template was modified
	LastModifiedDate *azuredevops.Time `json:"lastModifiedDate,omitempty"`
	// List of workspace mappings
	Mappings *[]WorkspaceMapping `json:"mappings,omitempty"`
	// Id of the workspace for this template
	WorkspaceId *int `json:"workspaceId,omitempty"`
}

type XamlBuildControllerReference struct {
	// Id of the resource
	Id *int `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

type XamlBuildDefinition struct {
	// The date this version of the definition was created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The ID of the referenced definition.
	Id *int `json:"id,omitempty"`
	// The name of the referenced definition.
	Name *string `json:"name,omitempty"`
	// The folder path of the definition.
	Path *string `json:"path,omitempty"`
	// A reference to the project.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// A value that indicates whether builds can be queued against this definition.
	QueueStatus *DefinitionQueueStatus `json:"queueStatus,omitempty"`
	// The definition revision number.
	Revision *int `json:"revision,omitempty"`
	// The type of the definition.
	Type *DefinitionType `json:"type,omitempty"`
	// The definition's URI.
	Uri *string `json:"uri,omitempty"`
	// The REST URL of the definition.
	Url   *string     `json:"url,omitempty"`
	Links interface{} `json:"_links,omitempty"`
	// Batch size of the definition
	BatchSize *int    `json:"batchSize,omitempty"`
	BuildArgs *string `json:"buildArgs,omitempty"`
	// The continuous integration quiet period
	ContinuousIntegrationQuietPeriod *int `json:"continuousIntegrationQuietPeriod,omitempty"`
	// The build controller
	Controller *BuildController `json:"controller,omitempty"`
	// The date this definition was created
	CreatedOn *azuredevops.Time `json:"createdOn,omitempty"`
	// Default drop location for builds from this definition
	DefaultDropLocation *string `json:"defaultDropLocation,omitempty"`
	// Description of the definition
	Description *string `json:"description,omitempty"`
	// The last build on this definition
	LastBuild *XamlBuildReference `json:"lastBuild,omitempty"`
	// The repository
	Repository *BuildRepository `json:"repository,omitempty"`
	// The reasons supported by the template
	SupportedReasons *BuildReason `json:"supportedReasons,omitempty"`
	// How builds are triggered from this definition
	TriggerType *DefinitionTriggerType `json:"triggerType,omitempty"`
}

type XamlBuildReference struct {
	// Id of the resource
	Id *int `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

type XamlBuildServerReference struct {
	// Id of the resource
	Id *int `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

type XamlDefinitionReference struct {
	// Id of the resource
	Id *int `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

// Represents a YAML process.
type YamlProcess struct {
	Errors *[]string `json:"errors,omitempty"`
	// The resources used by the build definition.
	Resources *BuildProcessResources `json:"resources,omitempty"`
	// The YAML filename.
	YamlFilename *string `json:"yamlFilename,omitempty"`
}
