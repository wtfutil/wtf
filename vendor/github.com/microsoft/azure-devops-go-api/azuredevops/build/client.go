// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package build

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var ResourceAreaId, _ = uuid.Parse("965220d5-5bb9-42cf-8d67-9b146df2a5a4")

type Client interface {
	// Adds a tag to a build.
	AddBuildTag(context.Context, AddBuildTagArgs) (*[]string, error)
	// Adds tags to a build.
	AddBuildTags(context.Context, AddBuildTagsArgs) (*[]string, error)
	// [Preview API] Adds a tag to a definition
	AddDefinitionTag(context.Context, AddDefinitionTagArgs) (*[]string, error)
	// [Preview API] Adds multiple tags to a definition.
	AddDefinitionTags(context.Context, AddDefinitionTagsArgs) (*[]string, error)
	// [Preview API]
	AuthorizeDefinitionResources(context.Context, AuthorizeDefinitionResourcesArgs) (*[]DefinitionResourceReference, error)
	// [Preview API]
	AuthorizeProjectResources(context.Context, AuthorizeProjectResourcesArgs) (*[]DefinitionResourceReference, error)
	// Associates an artifact with a build.
	CreateArtifact(context.Context, CreateArtifactArgs) (*BuildArtifact, error)
	// Creates a new definition.
	CreateDefinition(context.Context, CreateDefinitionArgs) (*BuildDefinition, error)
	// [Preview API] Creates a new folder.
	CreateFolder(context.Context, CreateFolderArgs) (*Folder, error)
	// Deletes a build.
	DeleteBuild(context.Context, DeleteBuildArgs) error
	// Removes a tag from a build.
	DeleteBuildTag(context.Context, DeleteBuildTagArgs) (*[]string, error)
	// Deletes a definition and all associated builds.
	DeleteDefinition(context.Context, DeleteDefinitionArgs) error
	// [Preview API] Removes a tag from a definition.
	DeleteDefinitionTag(context.Context, DeleteDefinitionTagArgs) (*[]string, error)
	// [Preview API] Deletes a definition folder. Definitions and their corresponding builds will also be deleted.
	DeleteFolder(context.Context, DeleteFolderArgs) error
	// Deletes a build definition template.
	DeleteTemplate(context.Context, DeleteTemplateArgs) error
	// Gets a specific artifact for a build.
	GetArtifact(context.Context, GetArtifactArgs) (*BuildArtifact, error)
	// Gets a specific artifact for a build.
	GetArtifactContentZip(context.Context, GetArtifactContentZipArgs) (io.ReadCloser, error)
	// Gets all artifacts for a build.
	GetArtifacts(context.Context, GetArtifactsArgs) (*[]BuildArtifact, error)
	// [Preview API] Gets a specific attachment.
	GetAttachment(context.Context, GetAttachmentArgs) (io.ReadCloser, error)
	// [Preview API] Gets the list of attachments of a specific type that are associated with a build.
	GetAttachments(context.Context, GetAttachmentsArgs) (*[]Attachment, error)
	// Gets a build
	GetBuild(context.Context, GetBuildArgs) (*Build, error)
	// [Preview API] Gets a badge that indicates the status of the most recent build for the specified branch.
	GetBuildBadge(context.Context, GetBuildBadgeArgs) (*BuildBadge, error)
	// [Preview API] Gets a badge that indicates the status of the most recent build for the specified branch.
	GetBuildBadgeData(context.Context, GetBuildBadgeDataArgs) (*string, error)
	// Gets the changes associated with a build
	GetBuildChanges(context.Context, GetBuildChangesArgs) (*GetBuildChangesResponseValue, error)
	// Gets a controller
	GetBuildController(context.Context, GetBuildControllerArgs) (*BuildController, error)
	// Gets controller, optionally filtered by name
	GetBuildControllers(context.Context, GetBuildControllersArgs) (*[]BuildController, error)
	// Gets an individual log file for a build.
	GetBuildLog(context.Context, GetBuildLogArgs) (io.ReadCloser, error)
	// Gets an individual log file for a build.
	GetBuildLogLines(context.Context, GetBuildLogLinesArgs) (*[]string, error)
	// Gets the logs for a build.
	GetBuildLogs(context.Context, GetBuildLogsArgs) (*[]BuildLog, error)
	// Gets the logs for a build.
	GetBuildLogsZip(context.Context, GetBuildLogsZipArgs) (io.ReadCloser, error)
	// Gets an individual log file for a build.
	GetBuildLogZip(context.Context, GetBuildLogZipArgs) (io.ReadCloser, error)
	// Gets all build definition options supported by the system.
	GetBuildOptionDefinitions(context.Context, GetBuildOptionDefinitionsArgs) (*[]BuildOptionDefinition, error)
	// [Preview API] Gets properties for a build.
	GetBuildProperties(context.Context, GetBuildPropertiesArgs) (interface{}, error)
	// [Preview API] Gets a build report.
	GetBuildReport(context.Context, GetBuildReportArgs) (*BuildReportMetadata, error)
	// [Preview API] Gets a build report.
	GetBuildReportHtmlContent(context.Context, GetBuildReportHtmlContentArgs) (io.ReadCloser, error)
	// Gets a list of builds.
	GetBuilds(context.Context, GetBuildsArgs) (*GetBuildsResponseValue, error)
	// Gets the build settings.
	GetBuildSettings(context.Context, GetBuildSettingsArgs) (*BuildSettings, error)
	// Gets the tags for a build.
	GetBuildTags(context.Context, GetBuildTagsArgs) (*[]string, error)
	// Gets details for a build
	GetBuildTimeline(context.Context, GetBuildTimelineArgs) (*Timeline, error)
	// Gets the work items associated with a build.
	GetBuildWorkItemsRefs(context.Context, GetBuildWorkItemsRefsArgs) (*[]webapi.ResourceRef, error)
	// Gets the work items associated with a build, filtered to specific commits.
	GetBuildWorkItemsRefsFromCommits(context.Context, GetBuildWorkItemsRefsFromCommitsArgs) (*[]webapi.ResourceRef, error)
	// [Preview API] Gets the changes made to the repository between two given builds.
	GetChangesBetweenBuilds(context.Context, GetChangesBetweenBuildsArgs) (*[]Change, error)
	// Gets a definition, optionally at a specific revision.
	GetDefinition(context.Context, GetDefinitionArgs) (*BuildDefinition, error)
	// [Preview API] Gets build metrics for a definition.
	GetDefinitionMetrics(context.Context, GetDefinitionMetricsArgs) (*[]BuildMetric, error)
	// [Preview API] Gets properties for a definition.
	GetDefinitionProperties(context.Context, GetDefinitionPropertiesArgs) (interface{}, error)
	// [Preview API]
	GetDefinitionResources(context.Context, GetDefinitionResourcesArgs) (*[]DefinitionResourceReference, error)
	// Gets all revisions of a definition.
	GetDefinitionRevisions(context.Context, GetDefinitionRevisionsArgs) (*[]BuildDefinitionRevision, error)
	// Gets a list of definitions.
	GetDefinitions(context.Context, GetDefinitionsArgs) (*GetDefinitionsResponseValue, error)
	// [Preview API] Gets the tags for a definition.
	GetDefinitionTags(context.Context, GetDefinitionTagsArgs) (*[]string, error)
	// Gets a file from the build.
	GetFile(context.Context, GetFileArgs) (io.ReadCloser, error)
	// [Preview API] Gets the contents of a file in the given source code repository.
	GetFileContents(context.Context, GetFileContentsArgs) (io.ReadCloser, error)
	// [Preview API] Gets a list of build definition folders.
	GetFolders(context.Context, GetFoldersArgs) (*[]Folder, error)
	// [Preview API] Gets the latest build for a definition, optionally scoped to a specific branch.
	GetLatestBuild(context.Context, GetLatestBuildArgs) (*Build, error)
	// [Preview API] Gets the contents of a directory in the given source code repository.
	GetPathContents(context.Context, GetPathContentsArgs) (*[]SourceRepositoryItem, error)
	// [Preview API] Gets build metrics for a project.
	GetProjectMetrics(context.Context, GetProjectMetricsArgs) (*[]BuildMetric, error)
	// [Preview API]
	GetProjectResources(context.Context, GetProjectResourcesArgs) (*[]DefinitionResourceReference, error)
	// [Preview API] Gets a pull request object from source provider.
	GetPullRequest(context.Context, GetPullRequestArgs) (*PullRequest, error)
	// [Preview API] Gets information about build resources in the system.
	GetResourceUsage(context.Context, GetResourceUsageArgs) (*BuildResourceUsage, error)
	// [Preview API] <p>Gets the build status for a definition, optionally scoped to a specific branch, stage, job, and configuration.</p> <p>If there are more than one, then it is required to pass in a stageName value when specifying a jobName, and the same rule then applies for both if passing a configuration parameter.</p>
	GetStatusBadge(context.Context, GetStatusBadgeArgs) (*string, error)
	// Gets a list of all build and definition tags in the project.
	GetTags(context.Context, GetTagsArgs) (*[]string, error)
	// Gets a specific build definition template.
	GetTemplate(context.Context, GetTemplateArgs) (*BuildDefinitionTemplate, error)
	// Gets all definition templates.
	GetTemplates(context.Context, GetTemplatesArgs) (*[]BuildDefinitionTemplate, error)
	// [Preview API] Gets all the work items between two builds.
	GetWorkItemsBetweenBuilds(context.Context, GetWorkItemsBetweenBuildsArgs) (*[]webapi.ResourceRef, error)
	// [Preview API] Gets a list of branches for the given source code repository.
	ListBranches(context.Context, ListBranchesArgs) (*[]string, error)
	// [Preview API] Gets a list of source code repositories.
	ListRepositories(context.Context, ListRepositoriesArgs) (*SourceRepositories, error)
	// [Preview API] Get a list of source providers and their capabilities.
	ListSourceProviders(context.Context, ListSourceProvidersArgs) (*[]SourceProviderAttributes, error)
	// [Preview API] Gets a list of webhooks installed in the given source code repository.
	ListWebhooks(context.Context, ListWebhooksArgs) (*[]RepositoryWebhook, error)
	// Queues a build
	QueueBuild(context.Context, QueueBuildArgs) (*Build, error)
	// Restores a deleted definition
	RestoreDefinition(context.Context, RestoreDefinitionArgs) (*BuildDefinition, error)
	// [Preview API] Recreates the webhooks for the specified triggers in the given source code repository.
	RestoreWebhooks(context.Context, RestoreWebhooksArgs) error
	// Updates an existing build definition template.
	SaveTemplate(context.Context, SaveTemplateArgs) (*BuildDefinitionTemplate, error)
	// Updates a build.
	UpdateBuild(context.Context, UpdateBuildArgs) (*Build, error)
	// [Preview API] Updates properties for a build.
	UpdateBuildProperties(context.Context, UpdateBuildPropertiesArgs) (interface{}, error)
	// Updates multiple builds.
	UpdateBuilds(context.Context, UpdateBuildsArgs) (*[]Build, error)
	// Updates the build settings.
	UpdateBuildSettings(context.Context, UpdateBuildSettingsArgs) (*BuildSettings, error)
	// Updates an existing definition.
	UpdateDefinition(context.Context, UpdateDefinitionArgs) (*BuildDefinition, error)
	// [Preview API] Updates properties for a definition.
	UpdateDefinitionProperties(context.Context, UpdateDefinitionPropertiesArgs) (interface{}, error)
	// [Preview API] Updates an existing folder at given  existing path
	UpdateFolder(context.Context, UpdateFolderArgs) (*Folder, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// Adds a tag to a build.
func (client *ClientImpl) AddBuildTag(ctx context.Context, args AddBuildTagArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.Tag == nil || *args.Tag == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Tag"}
	}
	routeValues["tag"] = *args.Tag

	locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddBuildTag function
type AddBuildTagArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The tag to add.
	Tag *string
}

// Adds tags to a build.
func (client *ClientImpl) AddBuildTags(ctx context.Context, args AddBuildTagsArgs) (*[]string, error) {
	if args.Tags == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Tags"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	body, marshalErr := json.Marshal(*args.Tags)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddBuildTags function
type AddBuildTagsArgs struct {
	// (required) The tags to add.
	Tags *[]string
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// [Preview API] Adds a tag to a definition
func (client *ClientImpl) AddDefinitionTag(ctx context.Context, args AddDefinitionTagArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)
	if args.Tag == nil || *args.Tag == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Tag"}
	}
	routeValues["tag"] = *args.Tag

	locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddDefinitionTag function
type AddDefinitionTagArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (required) The tag to add.
	Tag *string
}

// [Preview API] Adds multiple tags to a definition.
func (client *ClientImpl) AddDefinitionTags(ctx context.Context, args AddDefinitionTagsArgs) (*[]string, error) {
	if args.Tags == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Tags"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	body, marshalErr := json.Marshal(*args.Tags)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddDefinitionTags function
type AddDefinitionTagsArgs struct {
	// (required) The tags to add.
	Tags *[]string
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
}

// [Preview API]
func (client *ClientImpl) AuthorizeDefinitionResources(ctx context.Context, args AuthorizeDefinitionResourcesArgs) (*[]DefinitionResourceReference, error) {
	if args.Resources == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Resources"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	body, marshalErr := json.Marshal(*args.Resources)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ea623316-1967-45eb-89ab-e9e6110cf2d6")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DefinitionResourceReference
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AuthorizeDefinitionResources function
type AuthorizeDefinitionResourcesArgs struct {
	// (required)
	Resources *[]DefinitionResourceReference
	// (required) Project ID or project name
	Project *string
	// (required)
	DefinitionId *int
}

// [Preview API]
func (client *ClientImpl) AuthorizeProjectResources(ctx context.Context, args AuthorizeProjectResourcesArgs) (*[]DefinitionResourceReference, error) {
	if args.Resources == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Resources"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Resources)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("398c85bc-81aa-4822-947c-a194a05f0fef")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DefinitionResourceReference
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AuthorizeProjectResources function
type AuthorizeProjectResourcesArgs struct {
	// (required)
	Resources *[]DefinitionResourceReference
	// (required) Project ID or project name
	Project *string
}

// Associates an artifact with a build.
func (client *ClientImpl) CreateArtifact(ctx context.Context, args CreateArtifactArgs) (*BuildArtifact, error) {
	if args.Artifact == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Artifact"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	body, marshalErr := json.Marshal(*args.Artifact)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildArtifact
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateArtifact function
type CreateArtifactArgs struct {
	// (required) The artifact.
	Artifact *BuildArtifact
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// Creates a new definition.
func (client *ClientImpl) CreateDefinition(ctx context.Context, args CreateDefinitionArgs) (*BuildDefinition, error) {
	if args.Definition == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Definition"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.DefinitionToCloneId != nil {
		queryParams.Add("definitionToCloneId", strconv.Itoa(*args.DefinitionToCloneId))
	}
	if args.DefinitionToCloneRevision != nil {
		queryParams.Add("definitionToCloneRevision", strconv.Itoa(*args.DefinitionToCloneRevision))
	}
	body, marshalErr := json.Marshal(*args.Definition)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateDefinition function
type CreateDefinitionArgs struct {
	// (required) The definition.
	Definition *BuildDefinition
	// (required) Project ID or project name
	Project *string
	// (optional)
	DefinitionToCloneId *int
	// (optional)
	DefinitionToCloneRevision *int
}

// [Preview API] Creates a new folder.
func (client *ClientImpl) CreateFolder(ctx context.Context, args CreateFolderArgs) (*Folder, error) {
	if args.Folder == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Folder"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	body, marshalErr := json.Marshal(*args.Folder)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Folder
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateFolder function
type CreateFolderArgs struct {
	// (required) The folder.
	Folder *Folder
	// (required) Project ID or project name
	Project *string
	// (required) The full path of the folder.
	Path *string
}

// Deletes a build.
func (client *ClientImpl) DeleteBuild(ctx context.Context, args DeleteBuildArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteBuild function
type DeleteBuildArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// Removes a tag from a build.
func (client *ClientImpl) DeleteBuildTag(ctx context.Context, args DeleteBuildTagArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.Tag == nil || *args.Tag == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Tag"}
	}
	routeValues["tag"] = *args.Tag

	locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
	resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the DeleteBuildTag function
type DeleteBuildTagArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The tag to remove.
	Tag *string
}

// Deletes a definition and all associated builds.
func (client *ClientImpl) DeleteDefinition(ctx context.Context, args DeleteDefinitionArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteDefinition function
type DeleteDefinitionArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
}

// [Preview API] Removes a tag from a definition.
func (client *ClientImpl) DeleteDefinitionTag(ctx context.Context, args DeleteDefinitionTagArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)
	if args.Tag == nil || *args.Tag == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Tag"}
	}
	routeValues["tag"] = *args.Tag

	locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
	resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the DeleteDefinitionTag function
type DeleteDefinitionTagArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (required) The tag to remove.
	Tag *string
}

// [Preview API] Deletes a definition folder. Definitions and their corresponding builds will also be deleted.
func (client *ClientImpl) DeleteFolder(ctx context.Context, args DeleteFolderArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Path == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteFolder function
type DeleteFolderArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The full path to the folder.
	Path *string
}

// Deletes a build definition template.
func (client *ClientImpl) DeleteTemplate(ctx context.Context, args DeleteTemplateArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TemplateId == nil || *args.TemplateId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TemplateId"}
	}
	routeValues["templateId"] = *args.TemplateId

	locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTemplate function
type DeleteTemplateArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the template.
	TemplateId *string
}

// Gets a specific artifact for a build.
func (client *ClientImpl) GetArtifact(ctx context.Context, args GetArtifactArgs) (*BuildArtifact, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.ArtifactName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "artifactName"}
	}
	queryParams.Add("artifactName", *args.ArtifactName)
	locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildArtifact
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetArtifact function
type GetArtifactArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The name of the artifact.
	ArtifactName *string
}

// Gets a specific artifact for a build.
func (client *ClientImpl) GetArtifactContentZip(ctx context.Context, args GetArtifactContentZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.ArtifactName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "artifactName"}
	}
	queryParams.Add("artifactName", *args.ArtifactName)
	locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetArtifactContentZip function
type GetArtifactContentZipArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The name of the artifact.
	ArtifactName *string
}

// Gets all artifacts for a build.
func (client *ClientImpl) GetArtifacts(ctx context.Context, args GetArtifactsArgs) (*[]BuildArtifact, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildArtifact
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetArtifacts function
type GetArtifactsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// [Preview API] Gets a specific attachment.
func (client *ClientImpl) GetAttachment(ctx context.Context, args GetAttachmentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()
	if args.RecordId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RecordId"}
	}
	routeValues["recordId"] = (*args.RecordId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	locationId, _ := uuid.Parse("af5122d3-3438-485e-a25a-2dbbfde84ee6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetAttachment function
type GetAttachmentArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The ID of the timeline.
	TimelineId *uuid.UUID
	// (required) The ID of the timeline record.
	RecordId *uuid.UUID
	// (required) The type of the attachment.
	Type *string
	// (required) The name of the attachment.
	Name *string
}

// [Preview API] Gets the list of attachments of a specific type that are associated with a build.
func (client *ClientImpl) GetAttachments(ctx context.Context, args GetAttachmentsArgs) (*[]Attachment, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type

	locationId, _ := uuid.Parse("f2192269-89fa-4f94-baf6-8fb128c55159")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Attachment
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAttachments function
type GetAttachmentsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The type of attachment.
	Type *string
}

// Gets a build
func (client *ClientImpl) GetBuild(ctx context.Context, args GetBuildArgs) (*Build, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.PropertyFilters != nil {
		queryParams.Add("propertyFilters", *args.PropertyFilters)
	}
	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Build
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuild function
type GetBuildArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	BuildId *int
	// (optional)
	PropertyFilters *string
}

// [Preview API] Gets a badge that indicates the status of the most recent build for the specified branch.
func (client *ClientImpl) GetBuildBadge(ctx context.Context, args GetBuildBadgeArgs) (*BuildBadge, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepoType == nil || *args.RepoType == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepoType"}
	}
	routeValues["repoType"] = *args.RepoType

	queryParams := url.Values{}
	if args.RepoId != nil {
		queryParams.Add("repoId", *args.RepoId)
	}
	if args.BranchName != nil {
		queryParams.Add("branchName", *args.BranchName)
	}
	locationId, _ := uuid.Parse("21b3b9ce-fad5-4567-9ad0-80679794e003")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildBadge
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildBadge function
type GetBuildBadgeArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The repository type.
	RepoType *string
	// (optional) The repository ID.
	RepoId *string
	// (optional) The branch name.
	BranchName *string
}

// [Preview API] Gets a badge that indicates the status of the most recent build for the specified branch.
func (client *ClientImpl) GetBuildBadgeData(ctx context.Context, args GetBuildBadgeDataArgs) (*string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepoType == nil || *args.RepoType == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepoType"}
	}
	routeValues["repoType"] = *args.RepoType

	queryParams := url.Values{}
	if args.RepoId != nil {
		queryParams.Add("repoId", *args.RepoId)
	}
	if args.BranchName != nil {
		queryParams.Add("branchName", *args.BranchName)
	}
	locationId, _ := uuid.Parse("21b3b9ce-fad5-4567-9ad0-80679794e003")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue string
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildBadgeData function
type GetBuildBadgeDataArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The repository type.
	RepoType *string
	// (optional) The repository ID.
	RepoId *string
	// (optional) The branch name.
	BranchName *string
}

// Gets the changes associated with a build
func (client *ClientImpl) GetBuildChanges(ctx context.Context, args GetBuildChangesArgs) (*GetBuildChangesResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.IncludeSourceChange != nil {
		queryParams.Add("includeSourceChange", strconv.FormatBool(*args.IncludeSourceChange))
	}
	locationId, _ := uuid.Parse("54572c7b-bbd3-45d4-80dc-28be08941620")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetBuildChangesResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetBuildChanges function
type GetBuildChangesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	BuildId *int
	// (optional)
	ContinuationToken *string
	// (optional) The maximum number of changes to return
	Top *int
	// (optional)
	IncludeSourceChange *bool
}

// Return type for the GetBuildChanges function
type GetBuildChangesResponseValue struct {
	Value []Change
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// Gets a controller
func (client *ClientImpl) GetBuildController(ctx context.Context, args GetBuildControllerArgs) (*BuildController, error) {
	routeValues := make(map[string]string)
	if args.ControllerId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ControllerId"}
	}
	routeValues["controllerId"] = strconv.Itoa(*args.ControllerId)

	locationId, _ := uuid.Parse("fcac1932-2ee1-437f-9b6f-7f696be858f6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildController
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildController function
type GetBuildControllerArgs struct {
	// (required)
	ControllerId *int
}

// Gets controller, optionally filtered by name
func (client *ClientImpl) GetBuildControllers(ctx context.Context, args GetBuildControllersArgs) (*[]BuildController, error) {
	queryParams := url.Values{}
	if args.Name != nil {
		queryParams.Add("name", *args.Name)
	}
	locationId, _ := uuid.Parse("fcac1932-2ee1-437f-9b6f-7f696be858f6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildController
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildControllers function
type GetBuildControllersArgs struct {
	// (optional)
	Name *string
}

// Gets an individual log file for a build.
func (client *ClientImpl) GetBuildLog(ctx context.Context, args GetBuildLogArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	queryParams := url.Values{}
	if args.StartLine != nil {
		queryParams.Add("startLine", strconv.FormatUint(*args.StartLine, 10))
	}
	if args.EndLine != nil {
		queryParams.Add("endLine", strconv.FormatUint(*args.EndLine, 10))
	}
	locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBuildLog function
type GetBuildLogArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The ID of the log file.
	LogId *int
	// (optional) The start line.
	StartLine *uint64
	// (optional) The end line.
	EndLine *uint64
}

// Gets an individual log file for a build.
func (client *ClientImpl) GetBuildLogLines(ctx context.Context, args GetBuildLogLinesArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	queryParams := url.Values{}
	if args.StartLine != nil {
		queryParams.Add("startLine", strconv.FormatUint(*args.StartLine, 10))
	}
	if args.EndLine != nil {
		queryParams.Add("endLine", strconv.FormatUint(*args.EndLine, 10))
	}
	locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildLogLines function
type GetBuildLogLinesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The ID of the log file.
	LogId *int
	// (optional) The start line.
	StartLine *uint64
	// (optional) The end line.
	EndLine *uint64
}

// Gets the logs for a build.
func (client *ClientImpl) GetBuildLogs(ctx context.Context, args GetBuildLogsArgs) (*[]BuildLog, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildLog
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildLogs function
type GetBuildLogsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// Gets the logs for a build.
func (client *ClientImpl) GetBuildLogsZip(ctx context.Context, args GetBuildLogsZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBuildLogsZip function
type GetBuildLogsZipArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// Gets an individual log file for a build.
func (client *ClientImpl) GetBuildLogZip(ctx context.Context, args GetBuildLogZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	queryParams := url.Values{}
	if args.StartLine != nil {
		queryParams.Add("startLine", strconv.FormatUint(*args.StartLine, 10))
	}
	if args.EndLine != nil {
		queryParams.Add("endLine", strconv.FormatUint(*args.EndLine, 10))
	}
	locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBuildLogZip function
type GetBuildLogZipArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The ID of the log file.
	LogId *int
	// (optional) The start line.
	StartLine *uint64
	// (optional) The end line.
	EndLine *uint64
}

// Gets all build definition options supported by the system.
func (client *ClientImpl) GetBuildOptionDefinitions(ctx context.Context, args GetBuildOptionDefinitionsArgs) (*[]BuildOptionDefinition, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	locationId, _ := uuid.Parse("591cb5a4-2d46-4f3a-a697-5cd42b6bd332")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildOptionDefinition
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildOptionDefinitions function
type GetBuildOptionDefinitionsArgs struct {
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Gets properties for a build.
func (client *ClientImpl) GetBuildProperties(ctx context.Context, args GetBuildPropertiesArgs) (interface{}, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.Filter != nil {
		listAsString := strings.Join((*args.Filter)[:], ",")
		queryParams.Add("filter", listAsString)
	}
	locationId, _ := uuid.Parse("0a6312e9-0627-49b7-8083-7d74a64849c9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the GetBuildProperties function
type GetBuildPropertiesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (optional) A comma-delimited list of properties. If specified, filters to these specific properties.
	Filter *[]string
}

// [Preview API] Gets a build report.
func (client *ClientImpl) GetBuildReport(ctx context.Context, args GetBuildReportArgs) (*BuildReportMetadata, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.Type != nil {
		queryParams.Add("type", *args.Type)
	}
	locationId, _ := uuid.Parse("45bcaa88-67e1-4042-a035-56d3b4a7d44c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildReportMetadata
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildReport function
type GetBuildReportArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (optional)
	Type *string
}

// [Preview API] Gets a build report.
func (client *ClientImpl) GetBuildReportHtmlContent(ctx context.Context, args GetBuildReportHtmlContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.Type != nil {
		queryParams.Add("type", *args.Type)
	}
	locationId, _ := uuid.Parse("45bcaa88-67e1-4042-a035-56d3b4a7d44c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "text/html", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBuildReportHtmlContent function
type GetBuildReportHtmlContentArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (optional)
	Type *string
}

// Gets a list of builds.
func (client *ClientImpl) GetBuilds(ctx context.Context, args GetBuildsArgs) (*GetBuildsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Definitions != nil {
		var stringList []string
		for _, item := range *args.Definitions {
			stringList = append(stringList, strconv.Itoa(item))
		}
		listAsString := strings.Join((stringList)[:], ",")
		queryParams.Add("definitions", listAsString)
	}
	if args.Queues != nil {
		var stringList []string
		for _, item := range *args.Queues {
			stringList = append(stringList, strconv.Itoa(item))
		}
		listAsString := strings.Join((stringList)[:], ",")
		queryParams.Add("definitions", listAsString)
	}
	if args.BuildNumber != nil {
		queryParams.Add("buildNumber", *args.BuildNumber)
	}
	if args.MinTime != nil {
		queryParams.Add("minTime", (*args.MinTime).String())
	}
	if args.MaxTime != nil {
		queryParams.Add("maxTime", (*args.MaxTime).String())
	}
	if args.RequestedFor != nil {
		queryParams.Add("requestedFor", *args.RequestedFor)
	}
	if args.ReasonFilter != nil {
		queryParams.Add("reasonFilter", string(*args.ReasonFilter))
	}
	if args.StatusFilter != nil {
		queryParams.Add("statusFilter", string(*args.StatusFilter))
	}
	if args.ResultFilter != nil {
		queryParams.Add("resultFilter", string(*args.ResultFilter))
	}
	if args.TagFilters != nil {
		listAsString := strings.Join((*args.TagFilters)[:], ",")
		queryParams.Add("tagFilters", listAsString)
	}
	if args.Properties != nil {
		listAsString := strings.Join((*args.Properties)[:], ",")
		queryParams.Add("properties", listAsString)
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.MaxBuildsPerDefinition != nil {
		queryParams.Add("maxBuildsPerDefinition", strconv.Itoa(*args.MaxBuildsPerDefinition))
	}
	if args.DeletedFilter != nil {
		queryParams.Add("deletedFilter", string(*args.DeletedFilter))
	}
	if args.QueryOrder != nil {
		queryParams.Add("queryOrder", string(*args.QueryOrder))
	}
	if args.BranchName != nil {
		queryParams.Add("branchName", *args.BranchName)
	}
	if args.BuildIds != nil {
		var stringList []string
		for _, item := range *args.BuildIds {
			stringList = append(stringList, strconv.Itoa(item))
		}
		listAsString := strings.Join((stringList)[:], ",")
		queryParams.Add("definitions", listAsString)
	}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.RepositoryType != nil {
		queryParams.Add("repositoryType", *args.RepositoryType)
	}
	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetBuildsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetBuilds function
type GetBuildsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) A comma-delimited list of definition IDs. If specified, filters to builds for these definitions.
	Definitions *[]int
	// (optional) A comma-delimited list of queue IDs. If specified, filters to builds that ran against these queues.
	Queues *[]int
	// (optional) If specified, filters to builds that match this build number. Append * to do a prefix search.
	BuildNumber *string
	// (optional) If specified, filters to builds that finished/started/queued after this date based on the queryOrder specified.
	MinTime *azuredevops.Time
	// (optional) If specified, filters to builds that finished/started/queued before this date based on the queryOrder specified.
	MaxTime *azuredevops.Time
	// (optional) If specified, filters to builds requested for the specified user.
	RequestedFor *string
	// (optional) If specified, filters to builds that match this reason.
	ReasonFilter *BuildReason
	// (optional) If specified, filters to builds that match this status.
	StatusFilter *BuildStatus
	// (optional) If specified, filters to builds that match this result.
	ResultFilter *BuildResult
	// (optional) A comma-delimited list of tags. If specified, filters to builds that have the specified tags.
	TagFilters *[]string
	// (optional) A comma-delimited list of properties to retrieve.
	Properties *[]string
	// (optional) The maximum number of builds to return.
	Top *int
	// (optional) A continuation token, returned by a previous call to this method, that can be used to return the next set of builds.
	ContinuationToken *string
	// (optional) The maximum number of builds to return per definition.
	MaxBuildsPerDefinition *int
	// (optional) Indicates whether to exclude, include, or only return deleted builds.
	DeletedFilter *QueryDeletedOption
	// (optional) The order in which builds should be returned.
	QueryOrder *BuildQueryOrder
	// (optional) If specified, filters to builds that built branches that built this branch.
	BranchName *string
	// (optional) A comma-delimited list that specifies the IDs of builds to retrieve.
	BuildIds *[]int
	// (optional) If specified, filters to builds that built from this repository.
	RepositoryId *string
	// (optional) If specified, filters to builds that built from repositories of this type.
	RepositoryType *string
}

// Return type for the GetBuilds function
type GetBuildsResponseValue struct {
	Value []Build
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// Gets the build settings.
func (client *ClientImpl) GetBuildSettings(ctx context.Context, args GetBuildSettingsArgs) (*BuildSettings, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	locationId, _ := uuid.Parse("aa8c1c9c-ef8b-474a-b8c4-785c7b191d0d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildSettings
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildSettings function
type GetBuildSettingsArgs struct {
	// (optional) Project ID or project name
	Project *string
}

// Gets the tags for a build.
func (client *ClientImpl) GetBuildTags(ctx context.Context, args GetBuildTagsArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildTags function
type GetBuildTagsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// Gets details for a build
func (client *ClientImpl) GetBuildTimeline(ctx context.Context, args GetBuildTimelineArgs) (*Timeline, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)
	if args.TimelineId != nil {
		routeValues["timelineId"] = (*args.TimelineId).String()
	}

	queryParams := url.Values{}
	if args.ChangeId != nil {
		queryParams.Add("changeId", strconv.Itoa(*args.ChangeId))
	}
	if args.PlanId != nil {
		queryParams.Add("planId", (*args.PlanId).String())
	}
	locationId, _ := uuid.Parse("8baac422-4c6e-4de5-8532-db96d92acffa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Timeline
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildTimeline function
type GetBuildTimelineArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	BuildId *int
	// (optional)
	TimelineId *uuid.UUID
	// (optional)
	ChangeId *int
	// (optional)
	PlanId *uuid.UUID
}

// Gets the work items associated with a build.
func (client *ClientImpl) GetBuildWorkItemsRefs(ctx context.Context, args GetBuildWorkItemsRefsArgs) (*[]webapi.ResourceRef, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("5a21f5d2-5642-47e4-a0bd-1356e6731bee")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []webapi.ResourceRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildWorkItemsRefs function
type GetBuildWorkItemsRefsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (optional) The maximum number of work items to return.
	Top *int
}

// Gets the work items associated with a build, filtered to specific commits.
func (client *ClientImpl) GetBuildWorkItemsRefsFromCommits(ctx context.Context, args GetBuildWorkItemsRefsFromCommitsArgs) (*[]webapi.ResourceRef, error) {
	if args.CommitIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CommitIds"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	body, marshalErr := json.Marshal(*args.CommitIds)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("5a21f5d2-5642-47e4-a0bd-1356e6731bee")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []webapi.ResourceRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBuildWorkItemsRefsFromCommits function
type GetBuildWorkItemsRefsFromCommitsArgs struct {
	// (required) A comma-delimited list of commit IDs.
	CommitIds *[]string
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (optional) The maximum number of work items to return, or the number of commits to consider if no commit IDs are specified.
	Top *int
}

// [Preview API] Gets the changes made to the repository between two given builds.
func (client *ClientImpl) GetChangesBetweenBuilds(ctx context.Context, args GetChangesBetweenBuildsArgs) (*[]Change, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.FromBuildId != nil {
		queryParams.Add("fromBuildId", strconv.Itoa(*args.FromBuildId))
	}
	if args.ToBuildId != nil {
		queryParams.Add("toBuildId", strconv.Itoa(*args.ToBuildId))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("f10f0ea5-18a1-43ec-a8fb-2042c7be9b43")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Change
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetChangesBetweenBuilds function
type GetChangesBetweenBuildsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The ID of the first build.
	FromBuildId *int
	// (optional) The ID of the last build.
	ToBuildId *int
	// (optional) The maximum number of changes to return.
	Top *int
}

// Gets a definition, optionally at a specific revision.
func (client *ClientImpl) GetDefinition(ctx context.Context, args GetDefinitionArgs) (*BuildDefinition, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	queryParams := url.Values{}
	if args.Revision != nil {
		queryParams.Add("revision", strconv.Itoa(*args.Revision))
	}
	if args.MinMetricsTime != nil {
		queryParams.Add("minMetricsTime", (*args.MinMetricsTime).String())
	}
	if args.PropertyFilters != nil {
		listAsString := strings.Join((*args.PropertyFilters)[:], ",")
		queryParams.Add("propertyFilters", listAsString)
	}
	if args.IncludeLatestBuilds != nil {
		queryParams.Add("includeLatestBuilds", strconv.FormatBool(*args.IncludeLatestBuilds))
	}
	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDefinition function
type GetDefinitionArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (optional) The revision number to retrieve. If this is not specified, the latest version will be returned.
	Revision *int
	// (optional) If specified, indicates the date from which metrics should be included.
	MinMetricsTime *azuredevops.Time
	// (optional) A comma-delimited list of properties to include in the results.
	PropertyFilters *[]string
	// (optional)
	IncludeLatestBuilds *bool
}

// [Preview API] Gets build metrics for a definition.
func (client *ClientImpl) GetDefinitionMetrics(ctx context.Context, args GetDefinitionMetricsArgs) (*[]BuildMetric, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	queryParams := url.Values{}
	if args.MinMetricsTime != nil {
		queryParams.Add("minMetricsTime", (*args.MinMetricsTime).String())
	}
	locationId, _ := uuid.Parse("d973b939-0ce0-4fec-91d8-da3940fa1827")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildMetric
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDefinitionMetrics function
type GetDefinitionMetricsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (optional) The date from which to calculate metrics.
	MinMetricsTime *azuredevops.Time
}

// [Preview API] Gets properties for a definition.
func (client *ClientImpl) GetDefinitionProperties(ctx context.Context, args GetDefinitionPropertiesArgs) (interface{}, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	queryParams := url.Values{}
	if args.Filter != nil {
		listAsString := strings.Join((*args.Filter)[:], ",")
		queryParams.Add("filter", listAsString)
	}
	locationId, _ := uuid.Parse("d9826ad7-2a68-46a9-a6e9-677698777895")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the GetDefinitionProperties function
type GetDefinitionPropertiesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (optional) A comma-delimited list of properties. If specified, filters to these specific properties.
	Filter *[]string
}

// [Preview API]
func (client *ClientImpl) GetDefinitionResources(ctx context.Context, args GetDefinitionResourcesArgs) (*[]DefinitionResourceReference, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	locationId, _ := uuid.Parse("ea623316-1967-45eb-89ab-e9e6110cf2d6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DefinitionResourceReference
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDefinitionResources function
type GetDefinitionResourcesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	DefinitionId *int
}

// Gets all revisions of a definition.
func (client *ClientImpl) GetDefinitionRevisions(ctx context.Context, args GetDefinitionRevisionsArgs) (*[]BuildDefinitionRevision, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	locationId, _ := uuid.Parse("7c116775-52e5-453e-8c5d-914d9762d8c4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildDefinitionRevision
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDefinitionRevisions function
type GetDefinitionRevisionsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
}

// Gets a list of definitions.
func (client *ClientImpl) GetDefinitions(ctx context.Context, args GetDefinitionsArgs) (*GetDefinitionsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Name != nil {
		queryParams.Add("name", *args.Name)
	}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.RepositoryType != nil {
		queryParams.Add("repositoryType", *args.RepositoryType)
	}
	if args.QueryOrder != nil {
		queryParams.Add("queryOrder", string(*args.QueryOrder))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.MinMetricsTime != nil {
		queryParams.Add("minMetricsTime", (*args.MinMetricsTime).String())
	}
	if args.DefinitionIds != nil {
		var stringList []string
		for _, item := range *args.DefinitionIds {
			stringList = append(stringList, strconv.Itoa(item))
		}
		listAsString := strings.Join((stringList)[:], ",")
		queryParams.Add("definitions", listAsString)
	}
	if args.Path != nil {
		queryParams.Add("path", *args.Path)
	}
	if args.BuiltAfter != nil {
		queryParams.Add("builtAfter", (*args.BuiltAfter).String())
	}
	if args.NotBuiltAfter != nil {
		queryParams.Add("notBuiltAfter", (*args.NotBuiltAfter).String())
	}
	if args.IncludeAllProperties != nil {
		queryParams.Add("includeAllProperties", strconv.FormatBool(*args.IncludeAllProperties))
	}
	if args.IncludeLatestBuilds != nil {
		queryParams.Add("includeLatestBuilds", strconv.FormatBool(*args.IncludeLatestBuilds))
	}
	if args.TaskIdFilter != nil {
		queryParams.Add("taskIdFilter", (*args.TaskIdFilter).String())
	}
	if args.ProcessType != nil {
		queryParams.Add("processType", strconv.Itoa(*args.ProcessType))
	}
	if args.YamlFilename != nil {
		queryParams.Add("yamlFilename", *args.YamlFilename)
	}
	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetDefinitionsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetDefinitions function
type GetDefinitionsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) If specified, filters to definitions whose names match this pattern.
	Name *string
	// (optional) A repository ID. If specified, filters to definitions that use this repository.
	RepositoryId *string
	// (optional) If specified, filters to definitions that have a repository of this type.
	RepositoryType *string
	// (optional) Indicates the order in which definitions should be returned.
	QueryOrder *DefinitionQueryOrder
	// (optional) The maximum number of definitions to return.
	Top *int
	// (optional) A continuation token, returned by a previous call to this method, that can be used to return the next set of definitions.
	ContinuationToken *string
	// (optional) If specified, indicates the date from which metrics should be included.
	MinMetricsTime *azuredevops.Time
	// (optional) A comma-delimited list that specifies the IDs of definitions to retrieve.
	DefinitionIds *[]int
	// (optional) If specified, filters to definitions under this folder.
	Path *string
	// (optional) If specified, filters to definitions that have builds after this date.
	BuiltAfter *azuredevops.Time
	// (optional) If specified, filters to definitions that do not have builds after this date.
	NotBuiltAfter *azuredevops.Time
	// (optional) Indicates whether the full definitions should be returned. By default, shallow representations of the definitions are returned.
	IncludeAllProperties *bool
	// (optional) Indicates whether to return the latest and latest completed builds for this definition.
	IncludeLatestBuilds *bool
	// (optional) If specified, filters to definitions that use the specified task.
	TaskIdFilter *uuid.UUID
	// (optional) If specified, filters to definitions with the given process type.
	ProcessType *int
	// (optional) If specified, filters to YAML definitions that match the given filename.
	YamlFilename *string
}

// Return type for the GetDefinitions function
type GetDefinitionsResponseValue struct {
	Value []BuildDefinitionReference
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// [Preview API] Gets the tags for a definition.
func (client *ClientImpl) GetDefinitionTags(ctx context.Context, args GetDefinitionTagsArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	queryParams := url.Values{}
	if args.Revision != nil {
		queryParams.Add("revision", strconv.Itoa(*args.Revision))
	}
	locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDefinitionTags function
type GetDefinitionTagsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (optional) The definition revision number. If not specified, uses the latest revision of the definition.
	Revision *int
}

// Gets a file from the build.
func (client *ClientImpl) GetFile(ctx context.Context, args GetFileArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.ArtifactName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "artifactName"}
	}
	queryParams.Add("artifactName", *args.ArtifactName)
	if args.FileId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "fileId"}
	}
	queryParams.Add("fileId", *args.FileId)
	if args.FileName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "fileName"}
	}
	queryParams.Add("fileName", *args.FileName)
	locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetFile function
type GetFileArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (required) The name of the artifact.
	ArtifactName *string
	// (required) The primary key for the file.
	FileId *string
	// (required) The name that the file will be set to.
	FileName *string
}

// [Preview API] Gets the contents of a file in the given source code repository.
func (client *ClientImpl) GetFileContents(ctx context.Context, args GetFileContentsArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName

	queryParams := url.Values{}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	if args.Repository != nil {
		queryParams.Add("repository", *args.Repository)
	}
	if args.CommitOrBranch != nil {
		queryParams.Add("commitOrBranch", *args.CommitOrBranch)
	}
	if args.Path != nil {
		queryParams.Add("path", *args.Path)
	}
	locationId, _ := uuid.Parse("29d12225-b1d9-425f-b668-6c594a981313")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "text/plain", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetFileContents function
type GetFileContentsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
	// (optional) If specified, the vendor-specific identifier or the name of the repository to get branches. Can only be omitted for providers that do not support multiple repositories.
	Repository *string
	// (optional) The identifier of the commit or branch from which a file's contents are retrieved.
	CommitOrBranch *string
	// (optional) The path to the file to retrieve, relative to the root of the repository.
	Path *string
}

// [Preview API] Gets a list of build definition folders.
func (client *ClientImpl) GetFolders(ctx context.Context, args GetFoldersArgs) (*[]Folder, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.Path != nil && *args.Path != "" {
		routeValues["path"] = *args.Path
	}

	queryParams := url.Values{}
	if args.QueryOrder != nil {
		queryParams.Add("queryOrder", string(*args.QueryOrder))
	}
	locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Folder
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetFolders function
type GetFoldersArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The path to start with.
	Path *string
	// (optional) The order in which folders should be returned.
	QueryOrder *FolderQueryOrder
}

// [Preview API] Gets the latest build for a definition, optionally scoped to a specific branch.
func (client *ClientImpl) GetLatestBuild(ctx context.Context, args GetLatestBuildArgs) (*Build, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.Definition == nil || *args.Definition == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Definition"}
	}
	routeValues["definition"] = *args.Definition

	queryParams := url.Values{}
	if args.BranchName != nil {
		queryParams.Add("branchName", *args.BranchName)
	}
	locationId, _ := uuid.Parse("54481611-01f4-47f3-998f-160da0f0c229")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Build
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLatestBuild function
type GetLatestBuildArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) definition name with optional leading folder path, or the definition id
	Definition *string
	// (optional) optional parameter that indicates the specific branch to use
	BranchName *string
}

// [Preview API] Gets the contents of a directory in the given source code repository.
func (client *ClientImpl) GetPathContents(ctx context.Context, args GetPathContentsArgs) (*[]SourceRepositoryItem, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName

	queryParams := url.Values{}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	if args.Repository != nil {
		queryParams.Add("repository", *args.Repository)
	}
	if args.CommitOrBranch != nil {
		queryParams.Add("commitOrBranch", *args.CommitOrBranch)
	}
	if args.Path != nil {
		queryParams.Add("path", *args.Path)
	}
	locationId, _ := uuid.Parse("7944d6fb-df01-4709-920a-7a189aa34037")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []SourceRepositoryItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPathContents function
type GetPathContentsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
	// (optional) If specified, the vendor-specific identifier or the name of the repository to get branches. Can only be omitted for providers that do not support multiple repositories.
	Repository *string
	// (optional) The identifier of the commit or branch from which a file's contents are retrieved.
	CommitOrBranch *string
	// (optional) The path contents to list, relative to the root of the repository.
	Path *string
}

// [Preview API] Gets build metrics for a project.
func (client *ClientImpl) GetProjectMetrics(ctx context.Context, args GetProjectMetricsArgs) (*[]BuildMetric, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.MetricAggregationType != nil && *args.MetricAggregationType != "" {
		routeValues["metricAggregationType"] = *args.MetricAggregationType
	}

	queryParams := url.Values{}
	if args.MinMetricsTime != nil {
		queryParams.Add("minMetricsTime", (*args.MinMetricsTime).String())
	}
	locationId, _ := uuid.Parse("7433fae7-a6bc-41dc-a6e2-eef9005ce41a")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildMetric
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProjectMetrics function
type GetProjectMetricsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The aggregation type to use (hourly, daily).
	MetricAggregationType *string
	// (optional) The date from which to calculate metrics.
	MinMetricsTime *azuredevops.Time
}

// [Preview API]
func (client *ClientImpl) GetProjectResources(ctx context.Context, args GetProjectResourcesArgs) (*[]DefinitionResourceReference, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Type != nil {
		queryParams.Add("type", *args.Type)
	}
	if args.Id != nil {
		queryParams.Add("id", *args.Id)
	}
	locationId, _ := uuid.Parse("398c85bc-81aa-4822-947c-a194a05f0fef")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DefinitionResourceReference
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProjectResources function
type GetProjectResourcesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional)
	Type *string
	// (optional)
	Id *string
}

// [Preview API] Gets a pull request object from source provider.
func (client *ClientImpl) GetPullRequest(ctx context.Context, args GetPullRequestArgs) (*PullRequest, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName
	if args.PullRequestId == nil || *args.PullRequestId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = *args.PullRequestId

	queryParams := url.Values{}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	locationId, _ := uuid.Parse("d8763ec7-9ff0-4fb4-b2b2-9d757906ff14")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PullRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequest function
type GetPullRequestArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (required) Vendor-specific id of the pull request.
	PullRequestId *string
	// (optional) Vendor-specific identifier or the name of the repository that contains the pull request.
	RepositoryId *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
}

// [Preview API] Gets information about build resources in the system.
func (client *ClientImpl) GetResourceUsage(ctx context.Context, args GetResourceUsageArgs) (*BuildResourceUsage, error) {
	locationId, _ := uuid.Parse("3813d06c-9e36-4ea1-aac3-61a485d60e3d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildResourceUsage
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetResourceUsage function
type GetResourceUsageArgs struct {
}

// [Preview API] <p>Gets the build status for a definition, optionally scoped to a specific branch, stage, job, and configuration.</p> <p>If there are more than one, then it is required to pass in a stageName value when specifying a jobName, and the same rule then applies for both if passing a configuration parameter.</p>
func (client *ClientImpl) GetStatusBadge(ctx context.Context, args GetStatusBadgeArgs) (*string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.Definition == nil || *args.Definition == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Definition"}
	}
	routeValues["definition"] = *args.Definition

	queryParams := url.Values{}
	if args.BranchName != nil {
		queryParams.Add("branchName", *args.BranchName)
	}
	if args.StageName != nil {
		queryParams.Add("stageName", *args.StageName)
	}
	if args.JobName != nil {
		queryParams.Add("jobName", *args.JobName)
	}
	if args.Configuration != nil {
		queryParams.Add("configuration", *args.Configuration)
	}
	if args.Label != nil {
		queryParams.Add("label", *args.Label)
	}
	locationId, _ := uuid.Parse("07acfdce-4757-4439-b422-ddd13a2fcc10")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue string
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetStatusBadge function
type GetStatusBadgeArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Either the definition name with optional leading folder path, or the definition id.
	Definition *string
	// (optional) Only consider the most recent build for this branch.
	BranchName *string
	// (optional) Use this stage within the pipeline to render the status.
	StageName *string
	// (optional) Use this job within a stage of the pipeline to render the status.
	JobName *string
	// (optional) Use this job configuration to render the status
	Configuration *string
	// (optional) Replaces the default text on the left side of the badge.
	Label *string
}

// Gets a list of all build and definition tags in the project.
func (client *ClientImpl) GetTags(ctx context.Context, args GetTagsArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("d84ac5c6-edc7-43d5-adc9-1b34be5dea09")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTags function
type GetTagsArgs struct {
	// (required) Project ID or project name
	Project *string
}

// Gets a specific build definition template.
func (client *ClientImpl) GetTemplate(ctx context.Context, args GetTemplateArgs) (*BuildDefinitionTemplate, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TemplateId == nil || *args.TemplateId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TemplateId"}
	}
	routeValues["templateId"] = *args.TemplateId

	locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildDefinitionTemplate
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTemplate function
type GetTemplateArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the requested template.
	TemplateId *string
}

// Gets all definition templates.
func (client *ClientImpl) GetTemplates(ctx context.Context, args GetTemplatesArgs) (*[]BuildDefinitionTemplate, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BuildDefinitionTemplate
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTemplates function
type GetTemplatesArgs struct {
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Gets all the work items between two builds.
func (client *ClientImpl) GetWorkItemsBetweenBuilds(ctx context.Context, args GetWorkItemsBetweenBuildsArgs) (*[]webapi.ResourceRef, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.FromBuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "fromBuildId"}
	}
	queryParams.Add("fromBuildId", strconv.Itoa(*args.FromBuildId))
	if args.ToBuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "toBuildId"}
	}
	queryParams.Add("toBuildId", strconv.Itoa(*args.ToBuildId))
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("52ba8915-5518-42e3-a4bb-b0182d159e2d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []webapi.ResourceRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetWorkItemsBetweenBuilds function
type GetWorkItemsBetweenBuildsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the first build.
	FromBuildId *int
	// (required) The ID of the last build.
	ToBuildId *int
	// (optional) The maximum number of work items to return.
	Top *int
}

// [Preview API] Gets a list of branches for the given source code repository.
func (client *ClientImpl) ListBranches(ctx context.Context, args ListBranchesArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName

	queryParams := url.Values{}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	if args.Repository != nil {
		queryParams.Add("repository", *args.Repository)
	}
	if args.BranchName != nil {
		queryParams.Add("branchName", *args.BranchName)
	}
	locationId, _ := uuid.Parse("e05d4403-9b81-4244-8763-20fde28d1976")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListBranches function
type ListBranchesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
	// (optional) The vendor-specific identifier or the name of the repository to get branches. Can only be omitted for providers that do not support multiple repositories.
	Repository *string
	// (optional) If supplied, the name of the branch to check for specifically.
	BranchName *string
}

// [Preview API] Gets a list of source code repositories.
func (client *ClientImpl) ListRepositories(ctx context.Context, args ListRepositoriesArgs) (*SourceRepositories, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName

	queryParams := url.Values{}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	if args.Repository != nil {
		queryParams.Add("repository", *args.Repository)
	}
	if args.ResultSet != nil {
		queryParams.Add("resultSet", string(*args.ResultSet))
	}
	if args.PageResults != nil {
		queryParams.Add("pageResults", strconv.FormatBool(*args.PageResults))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("d44d1680-f978-4834-9b93-8c6e132329c9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue SourceRepositories
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListRepositories function
type ListRepositoriesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
	// (optional) If specified, the vendor-specific identifier or the name of a single repository to get.
	Repository *string
	// (optional) 'top' for the repositories most relevant for the endpoint. If not set, all repositories are returned. Ignored if 'repository' is set.
	ResultSet *ResultSet
	// (optional) If set to true, this will limit the set of results and will return a continuation token to continue the query.
	PageResults *bool
	// (optional) When paging results, this is a continuation token, returned by a previous call to this method, that can be used to return the next set of repositories.
	ContinuationToken *string
}

// [Preview API] Get a list of source providers and their capabilities.
func (client *ClientImpl) ListSourceProviders(ctx context.Context, args ListSourceProvidersArgs) (*[]SourceProviderAttributes, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("3ce81729-954f-423d-a581-9fea01d25186")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []SourceProviderAttributes
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListSourceProviders function
type ListSourceProvidersArgs struct {
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Gets a list of webhooks installed in the given source code repository.
func (client *ClientImpl) ListWebhooks(ctx context.Context, args ListWebhooksArgs) (*[]RepositoryWebhook, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName

	queryParams := url.Values{}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	if args.Repository != nil {
		queryParams.Add("repository", *args.Repository)
	}
	locationId, _ := uuid.Parse("8f20ff82-9498-4812-9f6e-9c01bdc50e99")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []RepositoryWebhook
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListWebhooks function
type ListWebhooksArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
	// (optional) If specified, the vendor-specific identifier or the name of the repository to get webhooks. Can only be omitted for providers that do not support multiple repositories.
	Repository *string
}

// Queues a build
func (client *ClientImpl) QueueBuild(ctx context.Context, args QueueBuildArgs) (*Build, error) {
	if args.Build == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Build"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.IgnoreWarnings != nil {
		queryParams.Add("ignoreWarnings", strconv.FormatBool(*args.IgnoreWarnings))
	}
	if args.CheckInTicket != nil {
		queryParams.Add("checkInTicket", *args.CheckInTicket)
	}
	if args.SourceBuildId != nil {
		queryParams.Add("sourceBuildId", strconv.Itoa(*args.SourceBuildId))
	}
	body, marshalErr := json.Marshal(*args.Build)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Build
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueueBuild function
type QueueBuildArgs struct {
	// (required)
	Build *Build
	// (required) Project ID or project name
	Project *string
	// (optional)
	IgnoreWarnings *bool
	// (optional)
	CheckInTicket *string
	// (optional)
	SourceBuildId *int
}

// Restores a deleted definition
func (client *ClientImpl) RestoreDefinition(ctx context.Context, args RestoreDefinitionArgs) (*BuildDefinition, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	queryParams := url.Values{}
	if args.Deleted == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "deleted"}
	}
	queryParams.Add("deleted", strconv.FormatBool(*args.Deleted))
	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the RestoreDefinition function
type RestoreDefinitionArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The identifier of the definition to restore.
	DefinitionId *int
	// (required) When false, restores a deleted definition.
	Deleted *bool
}

// [Preview API] Recreates the webhooks for the specified triggers in the given source code repository.
func (client *ClientImpl) RestoreWebhooks(ctx context.Context, args RestoreWebhooksArgs) error {
	if args.TriggerTypes == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.TriggerTypes"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.ProviderName == nil || *args.ProviderName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProviderName"}
	}
	routeValues["providerName"] = *args.ProviderName

	queryParams := url.Values{}
	if args.ServiceEndpointId != nil {
		queryParams.Add("serviceEndpointId", (*args.ServiceEndpointId).String())
	}
	if args.Repository != nil {
		queryParams.Add("repository", *args.Repository)
	}
	body, marshalErr := json.Marshal(*args.TriggerTypes)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("793bceb8-9736-4030-bd2f-fb3ce6d6b478")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RestoreWebhooks function
type RestoreWebhooksArgs struct {
	// (required) The types of triggers to restore webhooks for.
	TriggerTypes *[]DefinitionTriggerType
	// (required) Project ID or project name
	Project *string
	// (required) The name of the source provider.
	ProviderName *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
	ServiceEndpointId *uuid.UUID
	// (optional) If specified, the vendor-specific identifier or the name of the repository to get webhooks. Can only be omitted for providers that do not support multiple repositories.
	Repository *string
}

// Updates an existing build definition template.
func (client *ClientImpl) SaveTemplate(ctx context.Context, args SaveTemplateArgs) (*BuildDefinitionTemplate, error) {
	if args.Template == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Template"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TemplateId == nil || *args.TemplateId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TemplateId"}
	}
	routeValues["templateId"] = *args.TemplateId

	body, marshalErr := json.Marshal(*args.Template)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildDefinitionTemplate
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the SaveTemplate function
type SaveTemplateArgs struct {
	// (required) The new version of the template.
	Template *BuildDefinitionTemplate
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the template.
	TemplateId *string
}

// Updates a build.
func (client *ClientImpl) UpdateBuild(ctx context.Context, args UpdateBuildArgs) (*Build, error) {
	if args.Build == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Build"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	queryParams := url.Values{}
	if args.Retry != nil {
		queryParams.Add("retry", strconv.FormatBool(*args.Retry))
	}
	body, marshalErr := json.Marshal(*args.Build)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Build
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateBuild function
type UpdateBuildArgs struct {
	// (required) The build.
	Build *Build
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
	// (optional)
	Retry *bool
}

// [Preview API] Updates properties for a build.
func (client *ClientImpl) UpdateBuildProperties(ctx context.Context, args UpdateBuildPropertiesArgs) (interface{}, error) {
	if args.Document == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Document"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	body, marshalErr := json.Marshal(*args.Document)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0a6312e9-0627-49b7-8083-7d74a64849c9")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the UpdateBuildProperties function
type UpdateBuildPropertiesArgs struct {
	// (required) A json-patch document describing the properties to update.
	Document *[]webapi.JsonPatchOperation
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the build.
	BuildId *int
}

// Updates multiple builds.
func (client *ClientImpl) UpdateBuilds(ctx context.Context, args UpdateBuildsArgs) (*[]Build, error) {
	if args.Builds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Builds"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Builds)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Build
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateBuilds function
type UpdateBuildsArgs struct {
	// (required) The builds to update.
	Builds *[]Build
	// (required) Project ID or project name
	Project *string
}

// Updates the build settings.
func (client *ClientImpl) UpdateBuildSettings(ctx context.Context, args UpdateBuildSettingsArgs) (*BuildSettings, error) {
	if args.Settings == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Settings"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.Settings)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("aa8c1c9c-ef8b-474a-b8c4-785c7b191d0d")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildSettings
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateBuildSettings function
type UpdateBuildSettingsArgs struct {
	// (required) The new settings.
	Settings *BuildSettings
	// (optional) Project ID or project name
	Project *string
}

// Updates an existing definition.
func (client *ClientImpl) UpdateDefinition(ctx context.Context, args UpdateDefinitionArgs) (*BuildDefinition, error) {
	if args.Definition == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Definition"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	queryParams := url.Values{}
	if args.SecretsSourceDefinitionId != nil {
		queryParams.Add("secretsSourceDefinitionId", strconv.Itoa(*args.SecretsSourceDefinitionId))
	}
	if args.SecretsSourceDefinitionRevision != nil {
		queryParams.Add("secretsSourceDefinitionRevision", strconv.Itoa(*args.SecretsSourceDefinitionRevision))
	}
	body, marshalErr := json.Marshal(*args.Definition)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BuildDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateDefinition function
type UpdateDefinitionArgs struct {
	// (required) The new version of the definition.
	Definition *BuildDefinition
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
	// (optional)
	SecretsSourceDefinitionId *int
	// (optional)
	SecretsSourceDefinitionRevision *int
}

// [Preview API] Updates properties for a definition.
func (client *ClientImpl) UpdateDefinitionProperties(ctx context.Context, args UpdateDefinitionPropertiesArgs) (interface{}, error) {
	if args.Document == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Document"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DefinitionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DefinitionId"}
	}
	routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

	body, marshalErr := json.Marshal(*args.Document)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("d9826ad7-2a68-46a9-a6e9-677698777895")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the UpdateDefinitionProperties function
type UpdateDefinitionPropertiesArgs struct {
	// (required) A json-patch document describing the properties to update.
	Document *[]webapi.JsonPatchOperation
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the definition.
	DefinitionId *int
}

// [Preview API] Updates an existing folder at given  existing path
func (client *ClientImpl) UpdateFolder(ctx context.Context, args UpdateFolderArgs) (*Folder, error) {
	if args.Folder == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Folder"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	body, marshalErr := json.Marshal(*args.Folder)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Folder
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateFolder function
type UpdateFolderArgs struct {
	// (required) The new version of the folder.
	Folder *Folder
	// (required) Project ID or project name
	Project *string
	// (required) The full path to the folder.
	Path *string
}
